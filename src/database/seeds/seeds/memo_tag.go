package seeds

import (
	"fmt"

	"gorm.io/gorm"
)

type memo_tag struct {
	MemoID string
	TagID  string
}

func CreateMemoTag(db *gorm.DB, tid string, mid string) error {
	fmt.Println(db.Table("memo_tags").Create(&memo_tag{MemoID: mid, TagID: tid}).Error)
	return nil
}
func CreateMemoTags(db *gorm.DB, tids []string, mids []string) error {
	for i, v := range tids {
		if err := CreateMemoTag(db, v, mids[i]); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
