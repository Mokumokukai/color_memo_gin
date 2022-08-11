package presenter

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

//出力するデータの加工を行うインターフェースの定義
type memoPresenter struct {
}
type IColorMemoPresenter interface {
	ResponseColorMemos(memos []*models.ColorMemo) []*models.ColorMemo
}

func NewColorMemoPresenter() IColorMemoPresenter {
	return &memoPresenter{}
}

func (memoPresenter *memoPresenter) ResponseColorMemos(memos []*models.ColorMemo) []*models.ColorMemo {
	for _, v := range memos {
		v.ColorCode1 = "#" + v.ColorCode1
		v.ColorCode2 = "#" + v.ColorCode2
	}
	return memos
}
