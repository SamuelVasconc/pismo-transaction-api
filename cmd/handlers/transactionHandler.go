package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
	"github.com/SamuelVasconc/pismo-transaction-api/utils"
	"github.com/gin-gonic/gin"
)

// httpTransactionHandler represent the httphandler for Transaction
type httpTransactionHandler struct {
	transactionUseCase interfaces.TransactionUseCase
}

// NewTransactionHTTPHandler ...
func NewTransactionHTTPHandler(r *gin.RouterGroup, itransactionUseCase interfaces.TransactionUseCase) {
	handler := &httpTransactionHandler{
		transactionUseCase: itransactionUseCase,
	}
	r.POST("/transactions", handler.CreateNewTransaction)
}

// @Security Authorization
// @Summary List Transaction
// @Description
// @Accept  json
// @Produce  json
// @Param data body models.Account true "body request"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} models.Transaction
// @Router /accounts [post]
// @Router /accounts [get]
func (h *httpTransactionHandler) CreateNewTransaction(c *gin.Context) {

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[handlers/CreateNewTransaction] - Error on read parameters in request. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "")
		return
	}

	var outTransaction models.Transaction
	err = json.Unmarshal([]byte(payload), &outTransaction)
	if err != nil {
		log.Println("[handlers/CreateNewTransaction] - Error on parse parameters of request. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusBadRequest, "")
		return
	}

	newID, err := h.transactionUseCase.CreateNewTransaction(&outTransaction)

	if err == sql.ErrNoRows {
		utils.RespondWithError(c, http.StatusBadRequest, "This Operation Type does not exists.")
		return
	}

	if err != nil {
		log.Println("[handlers/CreateNewTransaction] - Error on save transaction: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "")
		return
	}

	utils.ResponseWithJSON(c, http.StatusOK, newID)
}
