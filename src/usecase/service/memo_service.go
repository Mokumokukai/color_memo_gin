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
	Duplicate(memo *models.ColorMemo) (*models.ColorMemo, error)
	Delete(memo *models.ColorMemo) error
	Edit(memo *models.ColorMemo) (*models.ColorMemo, error)
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

func (service *memoService) Duplicate(memo *models.ColorMemo) (*models.ColorMemo, error) {
	res, err := service.ColorMemoRepository.Duplicate(memo)
	if err != nil {
		return nil, err
	}
	return service.ColorMemoPresenter.ResponseColorMemo(res), nil
}

func (service *memoService) Delete(memo *models.ColorMemo) error {
	err := service.ColorMemoRepository.Delete(memo)
	return err
}
func (service *memoService) Edit(memo *models.ColorMemo) (*models.ColorMemo, error) {
	res, err := service.ColorMemoRepository.Edit(memo)
	if err != nil {
		return nil, err
	}
	return service.ColorMemoPresenter.ResponseColorMemo(res), nil
}
