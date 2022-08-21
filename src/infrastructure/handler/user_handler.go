package handler

import (
	"net/http"

	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/utils"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userController controllers.IUserController
}

type IUserHandler interface {
	GetUsers() gin.HandlerFunc
	Register() gin.HandlerFunc
}
type ReqSignup struct {
	Name string `json:"name"`
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
		req_u := &ReqSignup{}
		if err := c.Bind(req_u); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		var user models.User
		user.ID, _ = utils.AlphaNumNanoID(7)
		user.Name = req_u.Name
		uid, _ := c.Get("UID")
		user.UID, _ = uid.(string)
		u, err := handler.userController.Register(&user)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, u)
	}
}
