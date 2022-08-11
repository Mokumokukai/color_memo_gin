package service

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
)

type userService struct {
	UserRepository repository.IUserRepository
	UserPresenter  presenter.IUserPresenter
}

type IUserService interface {
	GetAll(users []*models.User) ([]*models.User, error)
}

func NewUserService(repository repository.IUserRepository, presentor presenter.IUserPresenter) IUserService {
	return &userService{repository, presentor}
}
func (service *userService) GetAll(users []*models.User) ([]*models.User, error) {
	res, err := service.UserRepository.GetAll(users)
	if err != nil {
		return nil, err
	}
	return service.UserPresenter.ResponseUsers(res), nil
}
