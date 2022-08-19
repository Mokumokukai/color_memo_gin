package handler

import (
	"net/http"

	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/Mokumokukai/color_memo_gin/src/utils"
	"github.com/gin-gonic/gin"
)

type tags_res struct {
	Tags []*models.Tag `json:"tags"`
}
type tag_res struct {
	Tag *models.Tag `json:"tag"`
}
type ReqTag struct {
	Tag *models.Tag `json:"tag"`
}
type tagHandler struct {
	tagController controllers.ITagController
}

type ITagHandler interface {
	GetTags() gin.HandlerFunc
	CreateTag() gin.HandlerFunc
}

func NewTagHandler(uc controllers.ITagController) ITagHandler {
	return &tagHandler{tagController: uc}
}

func (handler *tagHandler) GetTags() gin.HandlerFunc {

	return func(c *gin.Context) {
		t, err := handler.tagController.GetTags()
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, tags_res{Tags: t})
	}
}

func (handler *tagHandler) CreateTag() gin.HandlerFunc {

	return func(c *gin.Context) {
		req_t := &ReqTag{}

		if err := c.Bind(req_t); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		//validate here
		req_t.Tag.ID, _ = utils.AlphaNumNanoID(7)
		t, err := handler.tagController.CreateTag(req_t.Tag)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, tag_res{Tag: t})
	}
}
