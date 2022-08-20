package handler

import (
	"fmt"
	"net/http"

	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/validators"
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
	CreateTags() gin.HandlerFunc
}

func NewTagHandler(uc controllers.ITagController) ITagHandler {
	return &tagHandler{tagController: uc}
}

func (handler *tagHandler) GetTags() gin.HandlerFunc {

	return func(c *gin.Context) {
		ts, err := handler.tagController.GetTags()
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, tags_res{Tags: ts})
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
			return
		}
		c.JSON(200, tag_res{Tag: t})
	}
}
func (handler *tagHandler) CreateTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		b_bid := &validators.TagsValidator{}
		if err := c.Bind(b_bid); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		for _, v := range b_bid.Tags {
			fmt.Println(v)
		}
		ts, err := handler.tagController.CreateTags(b_bid.Tags)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(http.StatusCreated, tags_res{Tags: ts})

	}
}
