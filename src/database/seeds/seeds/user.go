package seeds

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name string, uid string) error {
	return db.Create(&models.User{ID: "aaaaaaaaa", UID: uid, Name: name}).Error
}
