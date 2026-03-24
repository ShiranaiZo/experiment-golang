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
	CreateUser(*gin.Context)
	GetUsers(*gin.Context)
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
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

		return
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

func (c *UserController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	user, err := c.service.GetUserService().Show(userId)

	code := http.StatusBadRequest
	message := http.StatusText(code)

	if err != nil {
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

func (c *UserController) UpdateUser(ctx *gin.Context) {
	data := dto.UserRequest{}
	userId := ctx.Param("userId")

	err := ctx.ShouldBindJSON(&data)
	code := http.StatusUnprocessableEntity
	message := http.StatusText(code)

	if err != nil {
		logrus.Errorf("Failed to bind JSON on UpdateUser controller: %v", err)

		helpers.HttpResponse(helpers.ParamHTTPResponse{
			Code:    code,
			Error:   &err,
			Message: &message,
			Ctx:     ctx,
		})

		return
	}

	user, err := c.service.GetUserService().UpdateUser(&data, userId)

	if err != nil {
		code = http.StatusBadRequest
		message = http.StatusText(code)
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

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	user, err := c.service.GetUserService().Show(userId)

	data := dto.UserRequest{
		Name:     user.Name,
		Address:  user.Address,
		Email:    user.Email,
		Password: user.Password,
	}

	code := http.StatusBadRequest
	message := http.StatusText(code)

	if err != nil {
		message = "User not found"
		logrus.Errorf("Failed to get user on DeleteUser controller: %v", err)
		helpers.HttpResponse(
			helpers.ParamHTTPResponse{
				Code:    code,
				Message: &message,
				Error:   &err,
				Ctx:     ctx,
			},
		)

		return
	}

	user, err = c.service.GetUserService().DeleteUser(&data, userId)

	if err != nil {
		message = "Failed to delete user "
		logrus.Errorf("Failed to delete user on DeleteUser controller: %v", err)
		helpers.HttpResponse(
			helpers.ParamHTTPResponse{
				Code:    code,
				Message: &message,
				Error:   &err,
				Ctx:     ctx,
			},
		)

		return
	}

	code = http.StatusOK
	message = http.StatusText(code)

	helpers.HttpResponse(
		helpers.ParamHTTPResponse{
			Code:    code,
			Message: &message,
			Ctx:     ctx,
		},
	)

}
