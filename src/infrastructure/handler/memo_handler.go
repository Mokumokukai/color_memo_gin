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
	DuplicateColorMemo() gin.HandlerFunc
	DeleteColorMemo() gin.HandlerFunc
	EditColorMemo() gin.HandlerFunc
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
			return
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
		user_id, _ := c.Get("user_id")
		req_m.ColorMemo.ID, _ = gonanoid.New(7)
		req_m.ColorMemo.CreaterID, _ = user_id.(string)
		req_m.ColorMemo.OwnerID = req_m.ColorMemo.CreaterID

		m, err := handler.memoController.CreateColorMemo(req_m.ColorMemo)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, memo_res{Memo: m})
	}
}

func (handler *memoHandler) DuplicateColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		memo := &models.ColorMemo{}
		//validate here
		user_id, _ := c.Get("user_id")
		memo.ID, _ = gonanoid.New(7)
		memo.OwnerID, _ = user_id.(string)

		m, err := handler.memoController.DuplicateColorMemo(c.Param("memo_id"), memo)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, memo_res{Memo: m})
	}
}

//Repositoryでidとowner_idを照会させるのでここでowner_idにリクエストを送ってきたuser_idを格納
func (handler *memoHandler) DeleteColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		memo := models.ColorMemo{}
		//validate here
		user_id, _ := c.Get("user_id")
		memo.OwnerID, _ = user_id.(string)
		memo.ID = c.Param("memo_id")

		err := handler.memoController.DeleteColorMemo(&memo)
		//TODO: errにエラー定義したエラーを入れるようにする
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusNoContent, "")
	}
}

//Repositoryでidとowner_idを照会させるのでここでowner_idにリクエストを送ってきたuser_idを格納
func (handler *memoHandler) EditColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		req_m := &ReqColorMemo{}
		if err := c.Bind(req_m); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		//validate here
		user_id, _ := c.Get("user_id")
		req_m.ColorMemo.ID = c.Param("memo_id")
		req_m.ColorMemo.OwnerID, _ = user_id.(string)

		m, err := handler.memoController.EditColorMemo(req_m.ColorMemo)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}
		c.JSON(http.StatusOK, memo_res{Memo: m})
	}
}
