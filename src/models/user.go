package models

type User struct {
	ID   string `json:"id"`
	UID  string `json:"-"`
	Name string `json:"name"`
}
