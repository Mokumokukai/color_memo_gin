package registry

import (
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/datastore"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/handler"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/service"

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

func (interactor *userInteractor) NewUserService() service.IUserService {
	return service.NewUserService(interactor.NewUserRepository(), interactor.NewUserPresenter())
}
func (interactor *userInteractor) NewUserRepository() repository.IUserRepository {
	return datastore.NewUserRepository(interactor.conn)
}

func (interactor *userInteractor) NewUserPresenter() presenter.IUserPresenter {
	return presenter.NewUserPresenter()
}
