package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	//factory
	_ "github.com/lib/pq"
)

//Server ...
type Server struct {
	Env string
}

//factory
var (
	DBConn *sql.DB
)

//InitDb ...
func InitDb() {

	a := Server{}
	a.Env = os.Getenv("ENV")
	connectionString := fmt.Sprintf("%s", a.GetDNS())
	var (
		err            error
		maxLifeTimeInt int
		maxIdleConns   int
		maxOpenConns   int
	)
	maxLifeTimeInt, _ = strconv.Atoi(os.Getenv("CONNMAXLIFETIME"))
	maxIdleConns, _ = strconv.Atoi(os.Getenv("MAXIDLECONNS"))
	maxOpenConns, _ = strconv.Atoi(os.Getenv("MAXOPENCONNS"))

	maxLifeTime := time.Duration(maxLifeTimeInt)

	DBConn, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}
	DBConn.SetConnMaxLifetime(time.Minute * maxLifeTime)
	DBConn.SetMaxIdleConns(maxIdleConns)
	DBConn.SetMaxOpenConns(maxOpenConns)

	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}
}

//GetDNS representa a recuperação do acesso ao banco
func (a *Server) GetDNS() string {
	var (
		dbUser     string
		dbPassword string
		dbname     string
		dbHost     string
		dbPort     string
	)

	dbUser = os.Getenv("DBUSER")
	dbPassword = os.Getenv("DBPASSWORD")
	dbname = os.Getenv("DBNAME")
	dbHost = os.Getenv("DBHOST")
	dbPort = os.Getenv("DBPORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbname)
	return connectionString
}
