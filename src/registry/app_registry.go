package registry

import (
	"gorm.io/gorm"
)

type interactor struct {
	userInteractor
}

type IInteractor interface {
	IUserInteractor
}

func NewInteractor(conn *gorm.DB) IInteractor {
	return &interactor{
		userInteractor{conn},
	}
}
