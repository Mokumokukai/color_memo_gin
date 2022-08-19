package valdators

type TagValidator struct {
	Tag struct {
		Name string `form:"name" json:"name" binding:"exists,alphanum,min=2,max=13"`
	}
}
