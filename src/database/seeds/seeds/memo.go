package seeds

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"

	"gorm.io/gorm"
)

func CreateColorMemo(db *gorm.DB, cid string, oid string, c1 string, c2 string) error {
	return db.Create(&models.ColorMemo{ID: "aaaaaaa", CreaterID: cid, OwnerID: oid, ColorCode1: c1, ColorCode2: c2}).Error
}
