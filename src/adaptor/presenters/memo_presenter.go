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
		v.ColorCode1 = "#" + v.ColorCode1
		v.ColorCode2 = "#" + v.ColorCode2
	}
	return memos
}
