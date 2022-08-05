package registry

import (
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/handler"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/repository"
	"github.com/Mokumokukai/color_memo_gin/src/usecase"

	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"gorm.io/gorm"
)

type userInteractor struct {
	conn *gorm.DB
}
type IUserInteractor interface {
	NewUserHandler() handler.IUserHandler
}

func NewUserInteractor(conn *gorm.DB) IUserInteractor {
	return &userInteractor{conn}
}

func (interactor *userInteractor) NewUserHandler() handler.IUserHandler {
	return handler.NewUserHandler(interactor.NewUserController())

}

func (interactor *userInteractor) NewUserController() controllers.IUserController {
	return controllers.NewUserController(interactor.NewUserService())
}

func (interactor *userInteractor) NewUserService() usecase.IUserService {
	return usecase.NewUserService(interactor.NewUserRepository(), interactor.NewUserPresenter())
}
func (interactor *userInteractor) NewUserRepository() usecase.IUserRepository {
	return repository.NewUserRepository(interactor.conn)
}

func (interactor *userInteractor) NewUserPresenter() usecase.IUserPresenter {
	return usecase.NewUserPresenter()
}
