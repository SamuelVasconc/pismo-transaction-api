package utils

import (
	"net/http"

	"github.com/SamuelVasconc/pismo-transaction-api/models"
	"github.com/gin-gonic/gin"
)

//ResponseWithJSON ...
func ResponseWithJSON(g *gin.Context, code int, payload interface{}) {
	g.JSON(code, payload)
}

//RespondWithError corresponde a funcao que restorna erro
func RespondWithError(g *gin.Context, code int, message string) {
	e := getMessageError(code, message)
	g.JSON(code, e)
}

//GetMessageError ...
func getMessageError(errorCode int, message string) *models.ResponseError {
	switch errorCode {
	case http.StatusInternalServerError:
		return &models.ResponseError{
			DeveloperMessage: "Internal server error",
			UserMessage:      "Was encountered an error when processing your request. We apologize for the inconvenience",
			MoreInfo:         "https://pismo.io/pt/",
			ErrorCode:        errorCode,
		}
	default:
		return &models.ResponseError{
			DeveloperMessage: "",
			UserMessage:      message,
			MoreInfo:         "https://pismo.io/pt/",
			ErrorCode:        404,
		}
	}
}

// EndpointNotFound ...
func EndpointNotFound(c *gin.Context) {
	c.Writer.WriteString("there's no endpoint for that.")
}
