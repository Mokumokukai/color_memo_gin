package presenters

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
)

type memoPresenter struct {
}

func NewColorMemoPresenter() presenter.IColorMemoPresenter {
	return &memoPresenter{}
}

func (memoPresenter *memoPresenter) ResponseColorMemos(memos []*models.ColorMemo) []*models.ColorMemo {
	for _, v := range memos {
		v.Color1 = "#" + v.Color1
		v.Color2 = "#" + v.Color2
	}
	return memos
}
