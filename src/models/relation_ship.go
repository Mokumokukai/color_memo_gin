package models

type RelationShip struct {
	ID          string `gorm:"primary_key"`
	FollowerID  string
	FollowingID string
}
