package seeds

import (
	"github.com/Mokumokukai/color_memo_gin/src/models"

	"gorm.io/gorm"
)

func CreateColorMemo(db *gorm.DB, id string, oid string, cid string, c1 string, c2 string) error {
	return db.Table("memos").Create(&models.ColorMemo{ID: id, OwnerID: oid, CreaterID: cid, Color1: c1, Color2: c2}).Error
}
