package validators

import "github.com/Mokumokukai/color_memo_gin/src/models"

type TagValidator struct {
	models.Tag
}
type TagsValidator struct {
	Tags []*models.Tag `form:"tags" json:"tags" binding:"dive" `
}
