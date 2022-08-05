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
}

func NewUserHandler(uc controllers.IUserController) IUserHandler {
	return &userHandler{userController: uc}
}

func (handler *userHandler) GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, "hello")
	}
}
