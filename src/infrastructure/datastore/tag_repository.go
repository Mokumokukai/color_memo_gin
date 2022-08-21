package datastore

import (
	"fmt"

	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
	"github.com/Mokumokukai/color_memo_gin/src/utils"
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
func (tagRepository *tagRepository) CreateMul(tags []*models.Tag) ([]*models.Tag, error) {
	return CreateTags(tagRepository.db, tags)
}

func CreateTags(db *gorm.DB, tags []*models.Tag) ([]*models.Tag, error) {
	//IDが存在しない場合はIDを挿入してtagを作成
	for _, tag := range tags {
		if tag.ID == "" {
			if err := db.Table("tags").Where(models.Tag{Name: tag.Name}).First(&tag).Error; err != nil {
				tag.ID, _ = utils.AlphaNumNanoID(7)
				db.Table("tags").Where(models.Tag{Name: tag.Name}).Create(&tag)
			}
		}
	}
	return tags, nil
}
