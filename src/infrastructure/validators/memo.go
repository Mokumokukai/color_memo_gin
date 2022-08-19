package validators

type ColorMemoValidator struct {
	ColorMemo struct {
		Color1 string `json:"color1" binding:"required,hexcolor"`
		Color2 string `json:"color2" binding:"required,hexcolor"`
	} `json:"memo"`
}
type MemoIDParams struct {
	MemoID string `uri:"memo_id" binding:"required,alphanum,len=7"`
}
