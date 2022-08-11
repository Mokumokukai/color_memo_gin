package presenter

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"
)

//出力するデータの加工を行うインターフェースの定義

type ITagPresenter interface {
	ResponseTags(tags []*models.Tag) []*models.Tag
	ResponseTag(tag *models.Tag) *models.Tag
}
