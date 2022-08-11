package models

type Tag struct {
	ID          string `gorm:"primary_key"`
	Name        string
	ColorMemoID string
}
