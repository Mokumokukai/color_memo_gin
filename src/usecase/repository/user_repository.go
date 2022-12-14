package repository

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type IUserRepository interface {
	GetAll(users []*models.User) ([]*models.User, error)
	Register(user *models.User) (*models.User, error)
}
