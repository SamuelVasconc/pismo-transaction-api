package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
	"github.com/SamuelVasconc/pismo-transaction-api/utils"
	"github.com/gin-gonic/gin"
)

// httpAccountHandler represent the httphandler for Account
type httpAccountHandler struct {
	accountUseCase interfaces.AccountUseCase
}

// NewAccountHTTPHandler ...
func NewAccountHTTPHandler(r *gin.RouterGroup, iaccountUseCase interfaces.AccountUseCase) {
	handler := &httpAccountHandler{
		accountUseCase: iaccountUseCase,
	}
	r.GET("/accounts/:accountId", handler.GetAccount)
	r.POST("/accounts", handler.CreateNewAccount)
}

// @Security Authorization
// @Summary List Accounts
// @Description
// @Accept  json
// @Produce  json
// @Param data body models.Account true "body request"
// @Success 200 {object} models.Account
// @Failure 400 {object} models.Account
// @Router /accounts [post]
// @Router /accounts [get]
func (h *httpAccountHandler) CreateNewAccount(c *gin.Context) {

	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[handlers/CreateNewAccount] - Error on read parameters in request. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "Error on read Account.")
		return
	}

	var outAccount models.Account
	err = json.Unmarshal(payload, &outAccount)
	if err != nil {
		log.Println("[handlers/CreateNewAccount] - Error on parse parameters of request. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusBadRequest, "Error on parse Account.")
		return
	}

	inAccount, err := h.accountUseCase.CreateNewAccount(&outAccount)
	if inAccount == nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		log.Println("[handlers/CreateNewAccount] - Error on create Account: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "")
		return
	}

	utils.ResponseWithJSON(c, http.StatusOK, inAccount)
}

func (h *httpAccountHandler) GetAccount(c *gin.Context) {

	outID := c.Param("accountId")
	convID, err := strconv.ParseInt(outID, 10, 64)
	if err != nil {
		log.Println("[handlers/GetAccount] - Error on parse parameters of request. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusBadRequest, "")
		return
	}

	account, err := h.accountUseCase.GetAccount(convID)
	if err != nil {
		log.Println("[handlers/GetAccount] - Error on get account. Erro: ", err.Error())
		utils.RespondWithError(c, http.StatusInternalServerError, "")
		return
	}

	if account == nil {
		utils.ResponseWithJSON(c, http.StatusNoContent, account)
	}

	utils.ResponseWithJSON(c, http.StatusOK, account)
}
