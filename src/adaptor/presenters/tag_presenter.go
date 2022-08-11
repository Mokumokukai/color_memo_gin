package presenters

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
)

type tagPresenter struct {
}

func NewTagPresenter() presenter.ITagPresenter {
	return &tagPresenter{}
}

func (tagPresenter *tagPresenter) ResponseTags(tags []*models.Tag) []*models.Tag {
	for _, v := range tags {
		v.Name = v.Name
	}
	return tags
}
func (tagPresenter *tagPresenter) ResponseTag(tag *models.Tag) *models.Tag {
	return tag
}
