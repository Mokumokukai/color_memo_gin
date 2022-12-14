package seeds

import (
	"fmt"

	"github.com/Mokumokukai/color_memo_gin/src/models"

	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, name string, id string) error {
	return db.Table("tags").Create(&models.Tag{Name: name, ID: id}).Error
}
func CreateTags(db *gorm.DB, names []string, ids []string) error {
	for i, v := range names {
		if err := CreateTag(db, v, ids[i]); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
