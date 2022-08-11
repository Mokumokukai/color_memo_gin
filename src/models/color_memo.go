package models

import "time"

type ColorMemo struct {
	ID        string    `json:"id" gorm:"primary_key"`
	OwnerID   string    `json:"owner_id"`
	CreaterID string    `json:"creater_id"`
	Color1    string    `json:"color1" binding:"required"`
	Color2    string    `json:"color2" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
