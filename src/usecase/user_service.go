package usecase

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type userService struct {
	UserRepository IUserRepository
	UserPresenter  IUserPresenter
}

type IUserService interface {
	GetAll(users []*models.User) ([]*models.User, error)
}

func NewUserService(repository IUserRepository, presentor IUserPresenter) IUserService {
	return &userService{repository, presentor}
}
func (service *userService) GetAll(users []*models.User) ([]*models.User, error) {
	res, err := service.UserRepository.GetAll(users)
	if err != nil {
		return nil, err
	}
	return service.UserPresenter.ResponseUsers(res), nil
}
