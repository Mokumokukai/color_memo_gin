package repository

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type IColorMemoRepository interface {
	GetAll(memos []*models.ColorMemo) ([]*models.ColorMemo, error)
	Create(memo *models.ColorMemo) (*models.ColorMemo, error)
	Duplicate(memo *models.ColorMemo) (*models.ColorMemo, error)
	Delete(memo *models.ColorMemo) error
	Edit(memo *models.ColorMemo) (*models.ColorMemo, error)
}
