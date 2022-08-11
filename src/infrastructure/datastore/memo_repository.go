package datastore

import (
	"fmt"

	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
	"gorm.io/gorm"
)

type memoRepository struct {
	db *gorm.DB
}

func NewColorMemoRepository(db *gorm.DB) repository.IColorMemoRepository {
	return &memoRepository{db}
}

func (memoRepository *memoRepository) GetAll(memos []*models.ColorMemo) ([]*models.ColorMemo, error) {
	err := memoRepository.db.Table("memos").Find(&memos).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return memos, err
}
