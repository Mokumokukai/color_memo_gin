package seeds

import (
	"gorm.io/gorm"
)

func All() []Seed {
	return []Seed{
		{
			Name: "CraeteTags",
			Run: func(db *gorm.DB) error {
				return CreateTags(db, []string{"test-1", "test-2", "test-3"}, []string{"mCKnkOg", "BYLsdos", "5sVSKa0"})
			},
		}, // {
		// 	Name: "CreateUser1",
		// 	Run: func(db *gorm.DB) error {
		// 		return CreateUser(db, "FzF3zyN", "test1", "uiduid1")
		// 	},
		// }, {
		// 	Name: "CreateUser2",
		// 	Run: func(db *gorm.DB) error {
		// 		return CreateUser(db, "Wl5c1Sk", "test2", "uiduid2")
		// 	},
		// }, {
		// 	Name: "CreateMemo1",
		// 	Run: func(db *gorm.DB) error {
		// 		return CreateColorMemo(db, "CQn_eQr", "Wl5c1Sk", "Wl5c1Sk", "00000000", "FFFFFFFF")
		// 	},
		// }, {
		// 	Name: "CreateMemo1",
		// 	Run: func(db *gorm.DB) error {
		// 		return CreateColorMemo(db, "813k2mA", "Wl5c1Sk", "Wl5c1Sk", "11000000", "FFFFFFFF")
		// 	},
		// },{
		{
			Name: "CraeteMemoTags",
			Run: func(db *gorm.DB) error {
				return CreateMemoTags(db, []string{"mCKnkOg", "mCKnkOg", "BYLsdos", "5sVSKa0"}, []string{"2Y4JTu9", "813k2mA", "813k2mA", "813k2mA"})

			},
		},
	}
}
