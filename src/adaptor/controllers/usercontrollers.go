package controllers

//依存性の注入
import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase"
)

type userController struct {
	userService usecase.IUserService
}

type IUserController interface {
	GetUsers() ([]*models.User, error)
}

func NewUserController(service usecase.IUserService) IUserController {
	return &userController{service}
}

func (userController *userController) GetUsers() ([]*models.User, error) {
	users := []*models.User{}
	res, err := userController.userService.GetAll(users)
	if err != nil {
		return nil, err
	}
	return res, nil
}
