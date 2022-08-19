package valdators

type ColorMemoValidator struct {
	ColorMemo struct {
		Color1 string `json:"color1" binding:"required,rgba"`
		Color2 string `json:"color2" binding:"required,rgba"`
	}
}
type MemoIDParams struct {
	MemoID string `url:"memo_id" binding:"required,alphanum,len=7"`
}
