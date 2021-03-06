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
//HealthCheck ...
func (h *httpHealthCheckHandler) HealthCheck(c *gin.Context) {

	err := db.DBConn.Ping()
	if err != nil {
		log.Println("[handlers/HealthCheck] - Erro when ping database. Error: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "")
		return
	}

	hc := &models.HealthCheck{DbUP: "UP", Status: "UP"}
	utils.ResponseWithJSON(c, http.StatusOK, hc)
}
