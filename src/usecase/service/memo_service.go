package service

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
)

type memoService struct {
	ColorMemoRepository repository.IColorMemoRepository
	ColorMemoPresenter  presenter.IColorMemoPresenter
}

type IColorMemoService interface {
	GetAll(memos []*models.ColorMemo) ([]*models.ColorMemo, error)
	Create(memo *models.ColorMemo) (*models.ColorMemo, error)
}

func NewColorMemoService(repository repository.IColorMemoRepository, presentor presenter.IColorMemoPresenter) IColorMemoService {
	return &memoService{repository, presentor}
}
func (service *memoService) GetAll(memos []*models.ColorMemo) ([]*models.ColorMemo, error) {
	res, err := service.ColorMemoRepository.GetAll(memos)
	if err != nil {
		return nil, err
	}
	return service.ColorMemoPresenter.ResponseColorMemos(res), nil
}
func (service *memoService) Create(memo *models.ColorMemo) (*models.ColorMemo, error) {
	res, err := service.ColorMemoRepository.Create(memo)
	if err != nil {
		return nil, err
	}
	return service.ColorMemoPresenter.ResponseColorMemo(res), nil
}
