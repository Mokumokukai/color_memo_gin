package seeds

import (
	"gorm.io/gorm"
)

func All() []Seed {
	return []Seed{
		{
			Name: "CreateUser1",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "FzF3zyN", "test1", "uiduid1")
			},
		}, {
			Name: "CreateUser2",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Wl5c1Sk", "test2", "uiduid2")
			},
		}, {
			Name: "CreateMemo1",
			Run: func(db *gorm.DB) error {
				return CreateColorMemo(db, "CQn_eQr", "Wl5c1Sk", "Wl5c1Sk", "00000000", "FFFFFFFF")
			},
		}, {
			Name: "CreateMemo1",
			Run: func(db *gorm.DB) error {
				return CreateColorMemo(db, "813k2mA", "Wl5c1Sk", "Wl5c1Sk", "11000000", "FFFFFFFF")
			},
		},
	}
}
