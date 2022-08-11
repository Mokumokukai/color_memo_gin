package registry

import (
	"gorm.io/gorm"
)

type interactor struct {
	userInteractor
	memoInteractor
}

type IInteractor interface {
	IUserInteractor
	IColorMemoInteractor
}

func NewInteractor(conn *gorm.DB) IInteractor {
	return &interactor{
		userInteractor{conn},
		memoInteractor{conn},
	}
}
