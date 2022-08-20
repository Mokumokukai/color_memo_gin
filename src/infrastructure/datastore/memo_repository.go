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
	err := memoRepository.db.Table("memos").Preload("Tags").Find(&memos).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return memos, nil
}

func (memoRepository *memoRepository) Create(memo *models.ColorMemo) (*models.ColorMemo, error) {
	err := memoRepository.db.Table("memos").Create(&memo).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	return memo, nil
}

// owner_idとidはそれぞれ複製した人のユーザIDと新規メモIDを挿入するので、一旦他の変数に保存する。
func (memoRepository *memoRepository) Duplicate(memo_id string, memo *models.ColorMemo) (*models.ColorMemo, error) {

	new_memo := &models.ColorMemo{ID: memo_id}
	if err := memoRepository.db.Table("memos").First(new_memo).Error; err != nil {
		return nil, fmt.Errorf("sql error", err)
	}
	new_memo.OwnerID = memo.OwnerID
	new_memo.ID = memo.ID
	if err := memoRepository.db.Table("memos").Create(new_memo).Error; err != nil {
		return nil, fmt.Errorf("sql error", err)

	}
	return new_memo, nil
}
func (memoRepository *memoRepository) Delete(memo *models.ColorMemo) error {

	if err := memoRepository.db.Table("memos").Where("id = ? AND owner_id = ?", memo.ID, memo.OwnerID).Delete(memo).Error; err != nil {
		return fmt.Errorf("sql error", err)

	}
	return nil
}

//Updatesは変更するものがなくてもエラーを返さないのでまず、存在確認を先に行う。
func (memoRepository *memoRepository) Edit(memo *models.ColorMemo) (*models.ColorMemo, error) {
	if err := memoRepository.db.Table("memos").Where("id = ? AND owner_id = ?", memo.ID, memo.OwnerID).Updates(memo).Error; err != nil {
		return nil, fmt.Errorf("cannot update", err)
	}
	return memo, nil
}
