package models

type Tag struct {
	ID   string `json:"id" gorm:"primary_key" binding:"len=0|len=7,alphanum|len=0"`
	Name string `json:"name" binding:"required,alphanum,min=2,max=13"`
}
