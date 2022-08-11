package models

type User struct {
	ID   string `json:"id" gorm:"primary_key"`
	UID  string `json:"-"`
	Name string `json:"name"`
}
