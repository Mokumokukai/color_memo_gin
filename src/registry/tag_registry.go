package registry

import (
	"github.com/Mokumokukai/color_memo_gin/src/adaptor/presenters"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/datastore"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/handler"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/presenter"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/repository"
	"github.com/Mokumokukai/color_memo_gin/src/usecase/service"

	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"gorm.io/gorm"
)

type tagInteractor struct {
	conn *gorm.DB
}
type ITagInteractor interface {
	NewTagHandler() handler.ITagHandler
}

func NewTagInteractor(conn *gorm.DB) ITagInteractor {
	return &tagInteractor{conn}
}

func (interactor *tagInteractor) NewTagHandler() handler.ITagHandler {
	return handler.NewTagHandler(interactor.NewTagController())

}

func (interactor *tagInteractor) NewTagController() controllers.ITagController {
	return controllers.NewTagController(interactor.NewTagService())
}

func (interactor *tagInteractor) NewTagService() service.ITagService {
	return service.NewTagService(interactor.NewTagRepository(), interactor.NewTagPresenter())
}
func (interactor *tagInteractor) NewTagRepository() repository.ITagRepository {
	return datastore.NewTagRepository(interactor.conn)
}

func (interactor *tagInteractor) NewTagPresenter() presenter.ITagPresenter {
	return presenters.NewTagPresenter()
}
