package handler

import (
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/Mokumokukai/color_memo_gin/src/models"
	"github.com/gin-gonic/gin"
)

type memos_res struct {
	Memos []*models.ColorMemo `json:"memos"`
}
type memo_res struct {
	Memo *models.ColorMemo `json:"memo"`
}
type memoHandler struct {
	memoController controllers.IColorMemoController
}

type IColorMemoHandler interface {
	GetColorMemos() gin.HandlerFunc
	CreateColorMemo() gin.HandlerFunc
}
type ReqColorMemo struct {
	ColorMemo *models.ColorMemo `json:"memo"`
}

func NewColorMemoHandler(mc controllers.IColorMemoController) IColorMemoHandler {
	return &memoHandler{memoController: mc}
}

func (handler *memoHandler) GetColorMemos() gin.HandlerFunc {

	return func(c *gin.Context) {
		m, err := handler.memoController.GetColorMemos()
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, memos_res{Memos: m})
	}
}

func (handler *memoHandler) CreateColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		req_m := &ReqColorMemo{}
		if err := c.Bind(req_m); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		//validate here
		req_m.ColorMemo.ID, _ = gonanoid.New(7)
		req_m.ColorMemo.OwnerID = req_m.ColorMemo.CreaterID
		m, err := handler.memoController.CreateColorMemo(req_m.ColorMemo)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, memo_res{Memo: m})
	}
}
