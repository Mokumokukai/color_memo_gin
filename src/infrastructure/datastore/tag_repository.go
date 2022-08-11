package datastore

import (
	"fmt"

	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) repository.ITagRepository {
	return &tagRepository{db}
}

func (tagRepository *tagRepository) GetAll(tags []*models.Tag) ([]*models.Tag, error) {
	err := tagRepository.db.Find(&tags).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return tags, err
}

func (tagRepository *tagRepository) Create(tag *models.Tag) (*models.Tag, error) {
	err := tagRepository.db.Table("tags").Create(&tag).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return tag, err
}
