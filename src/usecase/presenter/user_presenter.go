package presenter

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

//出力するデータの加工を行うインターフェースの定義
type userPresenter struct {
}
type IUserPresenter interface {
	ResponseUsers(users []*models.User) []*models.User
}

func NewUserPresenter() IUserPresenter {
	return &userPresenter{}
}

func (userPresenter *userPresenter) ResponseUsers(users []*models.User) []*models.User {
	for _, v := range users {
		v.Name = "!!!!" + v.Name
	}
	return users
}
