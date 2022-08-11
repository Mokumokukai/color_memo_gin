package presenter

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

//出力するデータの加工を行うインターフェースの定義

type IUserPresenter interface {
	ResponseUsers(users []*models.User) []*models.User
}
