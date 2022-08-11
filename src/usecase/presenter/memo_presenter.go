package presenter

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

type IColorMemoPresenter interface {
	ResponseColorMemos(memos []*models.ColorMemo) []*models.ColorMemo
}
