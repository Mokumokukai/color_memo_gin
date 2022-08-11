package controllers

//依存性の注入
import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/service"
)

type memoController struct {
	memoService service.IColorMemoService
}

type IColorMemoController interface {
	GetColorMemos() ([]*models.ColorMemo, error)
	CreateColorMemo(memo *models.ColorMemo) (*models.ColorMemo, error)
}

func NewColorMemoController(service service.IColorMemoService) IColorMemoController {
	return &memoController{service}
}

func (memoController *memoController) GetColorMemos() ([]*models.ColorMemo, error) {
	memos := []*models.ColorMemo{}
	res, err := memoController.memoService.GetAll(memos)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (memoController *memoController) CreateColorMemo(memo *models.ColorMemo) (*models.ColorMemo, error) {
	res, err := memoController.memoService.Create(memo)
	if err != nil {
		return nil, err
	}
	return res, nil
}
