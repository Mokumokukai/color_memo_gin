package seeds

import (
	"gorm.io/gorm"
)

func All() []Seed {
	return []Seed{
		{
			Name: "CreateUser1",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "test1", "uiduid1")
			},
		}, {
			Name: "CreateUser2",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "test2", "uiduid2")
			},
		},
	}
}
