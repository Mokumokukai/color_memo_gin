package presenters

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
)

type userPresenter struct {
}

func NewUserPresenter() presenter.IUserPresenter {
	return &userPresenter{}
}

func (userPresenter *userPresenter) ResponseUsers(users []*models.User) []*models.User {
	for _, v := range users {
		v.Name = "!!!!" + v.Name
	}
	return users
}
