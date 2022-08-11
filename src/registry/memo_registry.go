package registry

import (
	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/Mokumokukai/color_memo_gin/src/adaptor/presenters"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/datastore"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/handler"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/service"

	"gorm.io/gorm"
)

type memoInteractor struct {
	conn *gorm.DB
}
type IColorMemoInteractor interface {
	NewColorMemoHandler() handler.IColorMemoHandler
}

func NewColorMemoInteractor(conn *gorm.DB) IColorMemoInteractor {
	return &memoInteractor{conn}
}

func (interactor *memoInteractor) NewColorMemoHandler() handler.IColorMemoHandler {
	return handler.NewColorMemoHandler(interactor.NewColorMemoController())

}

func (interactor *memoInteractor) NewColorMemoController() controllers.IColorMemoController {
	return controllers.NewColorMemoController(interactor.NewColorMemoService())
}

func (interactor *memoInteractor) NewColorMemoService() service.IColorMemoService {
	return service.NewColorMemoService(interactor.NewColorMemoRepository(), interactor.NewColorMemoPresenter())
}
func (interactor *memoInteractor) NewColorMemoRepository() repository.IColorMemoRepository {
	return datastore.NewColorMemoRepository(interactor.conn)
}

func (interactor *memoInteractor) NewColorMemoPresenter() presenter.IColorMemoPresenter {
	return presenters.NewColorMemoPresenter()
}
