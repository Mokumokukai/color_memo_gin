package controllers

//依存性の注入
import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/service"
)

type tagController struct {
	tagService service.ITagService
}

type ITagController interface {
	GetTags() ([]*models.Tag, error)
	CreateTag(tag *models.Tag) (*models.Tag, error)
	CreateTags(tags []*models.Tag) ([]*models.Tag, error)
}

func NewTagController(service service.ITagService) ITagController {
	return &tagController{service}
}

func (tagController *tagController) GetTags() ([]*models.Tag, error) {
	tags := []*models.Tag{}
	res, err := tagController.tagService.GetAll(tags)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (tagController *tagController) CreateTag(tag *models.Tag) (*models.Tag, error) {
	res, err := tagController.tagService.Create(tag)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (tagController *tagController) CreateTags(tags []*models.Tag) ([]*models.Tag, error) {
	res, err := tagController.tagService.CreateMul(tags)
	if err != nil {
		return nil, err
	}
	return res, nil
}
