package service

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
)

type tagService struct {
	TagRepository repository.ITagRepository
	TagPresenter  presenter.ITagPresenter
}

type ITagService interface {
	GetAll(tags []*models.Tag) ([]*models.Tag, error)
	Create(tag *models.Tag) (*models.Tag, error)
}

func NewTagService(repository repository.ITagRepository, presentor presenter.ITagPresenter) ITagService {
	return &tagService{repository, presentor}
}
func (service *tagService) GetAll(tags []*models.Tag) ([]*models.Tag, error) {
	res, err := service.TagRepository.GetAll(tags)
	if err != nil {
		return nil, err
	}
	return service.TagPresenter.ResponseTags(res), nil
}
func (service *tagService) Create(tag *models.Tag) (*models.Tag, error) {
	res, err := service.TagRepository.Create(tag)
	if err != nil {
		return nil, err
	}
	return service.TagPresenter.ResponseTag(res), nil
}
