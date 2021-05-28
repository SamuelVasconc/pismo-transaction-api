package utils

import (
	"net/http"
	"reflect"

	"github.com/SamuelVasconc/pismo-transaction-api/models"
	"github.com/gin-gonic/gin"
)

//ResponseWithJSON ...
func ResponseWithJSON(g *gin.Context, code int, payload interface{}) {
	var r models.ResponseSuccess
	r.Records = payload
	lenPayload := reflect.ValueOf(payload)
	r.Meta.RecordCount = 1
	r.Meta.Limit = 1
	if lenPayload.Kind() == reflect.Slice {
		r.Meta.Limit = lenPayload.Len()
		r.Meta.RecordCount = lenPayload.Len()
	}
	g.JSON(code, r)
}

//RespondWithError corresponde a funcao que restorna erro
func RespondWithError(g *gin.Context, code int, message string, moreInfo string) {
	e := getMessageError(code)
	g.JSON(code, e)
}

//GetMessageError ...
func getMessageError(errorCode int) *models.ResponseError {
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
			DeveloperMessage: "Resource not found",
			UserMessage:      "Resource not found",
			MoreInfo:         "https://pismo.io/pt/",
			ErrorCode:        404,
		}
	}
}

// EndpointNotFound ...
func EndpointNotFound(c *gin.Context) {
	c.Writer.WriteString("there's no endpoint for that.")
}
