package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/SamuelVasconc/pismo-transaction-api/cmd/handlers"
	"github.com/SamuelVasconc/pismo-transaction-api/config/db"
	"github.com/SamuelVasconc/pismo-transaction-api/config/middlewares"
	"github.com/SamuelVasconc/pismo-transaction-api/docs"
	"github.com/SamuelVasconc/pismo-transaction-api/repositories"
	"github.com/SamuelVasconc/pismo-transaction-api/usecases"
	"github.com/SamuelVasconc/pismo-transaction-api/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Route            *gin.Engine
	RouteGroup       *gin.RouterGroup
	Port             string
	HttpReadTimeout  string
	HttpWriteTimeout string
	GroupRequest     string
}

func (s *Server) Initialization() {
	s.Port = os.Getenv("PORT")
	s.HttpReadTimeout = os.Getenv("HTTP_READ_TIMEOUT")
	s.HttpWriteTimeout = os.Getenv("HTTP_WRITE_TIMEOUT")
	s.GroupRequest = os.Getenv("GROUPREQUEST")

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Pismo Transaction API"
	docs.SwaggerInfo.Description = "API responsavel por cadastrar e manipular transações monetarias em contas. Case proposto pela Pismo."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = s.GroupRequest

	//Inicialize Database Connection
	db.InitDb()

	s.Route = gin.New()
	s.Route.Use(gin.Recovery())

	// Middleware
	m := new(middlewares.GoMiddleware)
	s.Route.Use(m.CORS())
	s.Route.NoRoute(utils.EndpointNotFound)
	s.RouteGroup = s.Route.Group(s.GroupRequest)

	//Account instances
	accountRepository := repositories.NewAccountRepository(db.DBConn)
	accountUseCase := usecases.NewAccountUseCase(accountRepository)

	//Transaction instances
	operationRepository := repositories.NewOperationRepository(db.DBConn)
	transactionRepository := repositories.NewTransactionRepository(db.DBConn)
	transactionUseCase := usecases.NewTransactionUseCase(transactionRepository, operationRepository, accountRepository)

	//Handler instances
	handlers.NewAccountHTTPHandler(s.RouteGroup, accountUseCase)
	handlers.NewTransactionHTTPHandler(s.RouteGroup, transactionUseCase)
	s.RouteGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func (s *Server) StartServer() {

	httpReadTimeout, _ := strconv.Atoi(s.HttpReadTimeout)
	httpWriteTimeout, _ := strconv.Atoi(s.HttpWriteTimeout)
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.Port),
		ReadTimeout:  time.Duration(httpReadTimeout) * time.Second,
		WriteTimeout: time.Duration(httpWriteTimeout) * time.Second,
		Handler:      s.Route,
	}

	log.Println("Starting Server on port: ", s.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
