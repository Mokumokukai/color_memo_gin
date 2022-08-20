package repository

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type ITagRepository interface {
	GetAll(tags []*models.Tag) ([]*models.Tag, error)
	Create(tag *models.Tag) (*models.Tag, error)
	CreateMul(tags []*models.Tag) ([]*models.Tag, error)
}
