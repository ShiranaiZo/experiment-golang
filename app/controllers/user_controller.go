package controllers

import (
	"github.com/ShiranaiZo/experiment-golang/app/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.IServiceRegistry
}

type IUserController interface {
	// List method controller
	// Login(*gin.Context)
	// Logout(*gin.Context)
	CreateUser(*gin.Context)
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

func (c *UserController) CreateUser(ctx *gin.Context) {

}

// func (c *UserController) GetUsers(ctx *gin.Context) {

// }

// func (c *UserController) GetUser(ctx *gin.Context, id string) {

// }

// func (c *UserController) EditUser(ctx *gin.Context, id string) {

// }

// func (c *UserController) DeleteUser(ctx *gin.Context, id string) {

// }
