package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParamHTTPResponse struct {
	Code    int
	Error   *error
	Message *string
	Data    interface{}
	Ctx     *gin.Context
	Token   *string
}

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Token   *string     `json:"token,omitempty"`
}

func HttpResponse(param ParamHTTPResponse) {
	var message interface{}

	message = http.StatusText(http.StatusInternalServerError)
	status := "error"

	if param.Error == nil {
		message = http.StatusText(http.StatusOK)
		status = "success"
	}

	if param.Message != nil {
		message = *param.Message
	}

	param.Ctx.JSON(param.Code, Response{
		Status:  status,
		Message: message,
		Data:    param.Data,
		Token:   param.Token,
	})
}
