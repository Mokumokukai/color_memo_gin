package seeds

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, id string, uid string, name string) error {
	return db.Create(&models.User{ID: id, UID: uid, Name: name}).Error
}
