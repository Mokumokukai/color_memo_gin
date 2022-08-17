package handler

import (
	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userController controllers.IUserController
}

type IUserHandler interface {
	GetUsers() gin.HandlerFunc
	Register() gin.HandlerFunc
}

func NewUserHandler(uc controllers.IUserController) IUserHandler {
	return &userHandler{userController: uc}
}

func (handler *userHandler) GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		m, err := handler.userController.GetUsers()
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, m)
	}
}

func (handler *userHandler) Register() gin.HandlerFunc {

	return func(c *gin.Context) {
		m, err := handler.userController.Register()
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, m)
	}
}
