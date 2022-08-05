package repository

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usecase.IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) GetAll(users []*models.User) ([]*models.User, error) {
	err := userRepository.db.Find(&users).Error

	return users, err
}
