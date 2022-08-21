package seeds

import (
	"gorm.io/gorm"
)

var tags_ids = []string{"mgDCzCq", "a5MKnMN", "im9ewKp", "JaNaWUT", "2YtobQt", "IZazuDB", "dat5j9b", "ysrrlJt", "WrwgyPV", "wi5cxUI"}
var memos_ids = []string{"WPSjSRs", "ay7zALC"}
var tag_names = []string{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test9", "test10"}
var user_ids = []string{"Qeeqd4E", "1oxf9vG"}

func All() []Seed {
	return []Seed{
		{
			Name: "CraeteTags",
			Run: func(db *gorm.DB) error {
				return CreateTags(db, tag_names, tags_ids)
			},
		}, {
			Name: "CreateUser1",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, user_ids[0], "user1", "uiduid1")
			},
		}, {
			Name: "CreateUser2",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, user_ids[1], "user2", "uiduid2")
			},
		}, {
			Name: "CreateMemo1",
			Run: func(db *gorm.DB) error {
				return CreateColorMemo(db, memos_ids[0], user_ids[0], user_ids[0], "00000000", "FFFFFFFF")
			},
		}, {
			Name: "CreateMemo1",
			Run: func(db *gorm.DB) error {
				return CreateColorMemo(db, memos_ids[1], user_ids[1], user_ids[1], "11000000", "FFFFFFFF")
			},
		}, {
			Name: "CraeteMemoTags",
			Run: func(db *gorm.DB) error {
				return CreateMemoTags(db, []string{tags_ids[0], tags_ids[1], tags_ids[2], tags_ids[0], tags_ids[1]}, []string{memos_ids[0], memos_ids[0], memos_ids[0], memos_ids[1], memos_ids[1]})

			},
		},
	}
}
