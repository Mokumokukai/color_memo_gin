package validators

import "github.com/Mokumokukai/color_memo_gin/src/models"

type ColorMemoValidator struct {
	ColorMemo models.ColorMemo `json:"memo"`
}
type MemoIDParams struct {
	MemoID string `uri:"memo_id" binding:"required,alphanum,len=7"`
}
