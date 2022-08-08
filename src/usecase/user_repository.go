package usecase

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type IUserRepository interface {
	GetAll(users []*models.User) ([]*models.User, error)
}
