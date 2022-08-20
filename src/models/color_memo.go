package models

import "time"

type ColorMemo struct {
	ID        string    `json:"id" gorm:"primary_key"`
	OwnerID   string    `json:"owner_id"`
	CreaterID string    `json:"creater_id"`
	Color1    string    `json:"color1" binding:"required,hexcolor"`
	Color2    string    `json:"color2" binding:"required,hexcolor"`
	CreatedAt time.Time `json:"created_at" gorm:"->"`
	UpdatedAt time.Time `json:"updated_at" gorm:"->"`
	Tags      []*Tag    `json:"tags" gorm:"many2many:memo_tags;foreignKey:ID;joinForeignKey:MemoID;References:ID;joinReferences:TagID" binding:"dive"`
}
