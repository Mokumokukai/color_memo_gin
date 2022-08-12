package registry

import (
	"gorm.io/gorm"
)

type interactor struct {
	userInteractor
	memoInteractor
	tagInteractor
}

type IInteractor interface {
	IUserInteractor
	IColorMemoInteractor
	ITagInteractor
}

func NewInteractor(conn *gorm.DB) IInteractor {
	return &interactor{
		userInteractor{conn},
		memoInteractor{conn},
		tagInteractor{conn},
	}
}
