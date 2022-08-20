package validators

import "github.com/Mokumokukai/color_memo_gin/src/models"

type ColorMemoValidator struct {
	ColorMemo struct {
		Color1 string        `json:"color1" binding:"required,hexcolor"`
		Color2 string        `json:"color2" binding:"required,hexcolor"`
		Tags   []*models.Tag `form:"tags" json:"tags" binding:"dive" `
	} `json:"memo"`
}
type MemoIDParams struct {
	MemoID string `uri:"memo_id" binding:"required,alphanum,len=7"`
}
