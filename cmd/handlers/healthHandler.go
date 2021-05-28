package handlers

import (
	"log"
	"net/http"

	"github.com/SamuelVasconc/pismo-transaction-api/models"

	"github.com/SamuelVasconc/pismo-transaction-api/config/db"
	"github.com/SamuelVasconc/pismo-transaction-api/utils"

	"github.com/gin-gonic/gin"
)

// httpHealthCheckHandler represent the httphandler for healthcheck
type httpHealthCheckHandler struct {
}

// NewHealthCheckHTTPHandler ...
func NewHealthCheckHTTPHandler(r *gin.RouterGroup) {
	handler := &httpHealthCheckHandler{}
	r.GET("/health", handler.HealthCheck)
}

// @Summary HealthCheck
// @Description HealthCheck API
// @Failure 400 {object} models.ResponseError
// @Success 200 {object} models.HealthCheck
// @Router /health [get]
func (h *httpHealthCheckHandler) HealthCheck(c *gin.Context) {

	err := db.DBConn.Ping()
	if err != nil {
		log.Println("[handlers/HealthCheck] - Erro ao realizar ping no banco de dados. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "No Connection with DataBase", err.Error())
		return
	}

	hc := &models.HealthCheck{DbUP: "UP", Status: "UP"}
	utils.ResponseWithJSON(c, http.StatusOK, hc)
}
