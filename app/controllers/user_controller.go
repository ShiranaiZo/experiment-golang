package controllers

import (
	"net/http"

	"github.com/ShiranaiZo/experiment-golang/app/database/dto"
	"github.com/ShiranaiZo/experiment-golang/app/services"
	"github.com/ShiranaiZo/experiment-golang/helpers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	service services.IServiceRegistry
}

type IUserController interface {
	// List method controller
	// Login(*gin.Context)
	// Logout(*gin.Context)
	CreateUser(*gin.Context)
	GetUsers(*gin.Context)
	// GetUsers(*gin.Context)
	// GetUser(*gin.Context, string)
	// EditUser(*gin.Context, string)
	// DeleteUser(*gin.Context, string)
}

func NewUserController(service services.IServiceRegistry) IUserController {
	return &UserController{
		service: service,
	}
}

func (c UserController) CreateUser(ctx *gin.Context) {
	data := dto.UserRequest{}

	err := ctx.ShouldBindJSON(&data)
	code := http.StatusUnprocessableEntity
	message := http.StatusText(code)

	if err != nil {
		logrus.Errorf("Failed to bind JSON on CreateUser controller: %v", err)
		helpers.HttpResponse(helpers.ParamHTTPResponse{
			Code:    code,
			Error:   &err,
			Message: &message,
			Ctx:     ctx,
		})

		return
	}
	user, err := c.service.GetUserService().Create(&data)

	if err != nil {
		logrus.Errorf("Failed to create user on CreateUser controller: %v", err)

		code := http.StatusBadRequest
		message := http.StatusText(code)

		helpers.HttpResponse(helpers.ParamHTTPResponse{
			Code:    code,
			Error:   &err,
			Message: &message,
			Ctx:     ctx,
		})

		return
	}

	code = http.StatusOK
	message = http.StatusText(code)

	helpers.HttpResponse(helpers.ParamHTTPResponse{
		Code:    code,
		Message: &message,
		Data:    user,
		Ctx:     ctx,
	})
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.service.GetUserService().Index()

	code := http.StatusBadRequest
	message := http.StatusText(code)

	if err != nil {
		helpers.HttpResponse(
			helpers.ParamHTTPResponse{
				Code:    code,
				Error:   &err,
				Message: &message,
				Ctx:     ctx,
			},
		)
	}

	code = http.StatusOK
	message = http.StatusText(code)

	helpers.HttpResponse(
		helpers.ParamHTTPResponse{
			Code:    code,
			Message: &message,
			Data:    users,
			Ctx:     ctx,
		},
	)
}

// func (c *UserController) GetUser(ctx *gin.Context, id string) {

// }

// func (c *UserController) EditUser(ctx *gin.Context, id string) {

// }

// func (c *UserController) DeleteUser(ctx *gin.Context, id string) {

// }
