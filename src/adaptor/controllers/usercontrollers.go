package controllers

//依存性の注入
import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/service"
)

type userController struct {
	userService service.IUserService
}

type IUserController interface {
	GetUsers() ([]*models.User, error)
	Register() (*models.User, error)
}

func NewUserController(service service.IUserService) IUserController {
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

func (userController *userController) Register() (*models.User, error) {
	user := *&models.User{}
	res, err := userController.userService.Register(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
