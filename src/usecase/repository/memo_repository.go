package repository

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type IColorMemoRepository interface {
	GetAll(memos []*models.ColorMemo) ([]*models.ColorMemo, error)
}
