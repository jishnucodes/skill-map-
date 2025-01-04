package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestResponse struct {
	Message string      `json:"message"`
	Status  uint        `json:"status"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, msg string, data interface{}) {
	response := requestResponse{
		msg,
		http.StatusOK,
		data,
	}
	ctx.JSON(http.StatusOK, response)
}

func BadResponse(ctx *gin.Context, msg string, data interface{}) {
	response := requestResponse{
		msg,
		http.StatusBadRequest,
		data,
	}
	ctx.JSON(http.StatusBadRequest, response)
}