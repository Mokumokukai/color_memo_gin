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
		//body部分のバリデーション
		b_bid := &validators.ColorMemoValidator{}
		if err := c.Bind(b_bid); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		user_id, _ := c.Get("user_id")
		id, _ := utils.AlphaNumNanoID(7)
		creater_id, _ := user_id.(string)
		var color_memo = &models.ColorMemo{
			ID:        id,
			CreaterID: creater_id,
			OwnerID:   creater_id,
			Color1:    b_bid.ColorMemo.Color1[1:],
			Color2:    b_bid.ColorMemo.Color2[1:],
		}
		fmt.Println(color_memo)
		m, err := handler.memoController.CreateColorMemo(color_memo)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, memo_res{Memo: m})
	}
}

func (handler *memoHandler) DuplicateColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		//uriをバリデーション :memo_idにalphanum以外が入らないように
		uri_bid := &validators.MemoIDParams{}
		if err := c.ShouldBindUri(uri_bid); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		memo := &models.ColorMemo{}
		user_id, _ := c.Get("user_id")
		memo.ID, _ = utils.AlphaNumNanoID(7)
		memo.OwnerID, _ = user_id.(string)

		m, err := handler.memoController.DuplicateColorMemo(uri_bid.MemoID, memo)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, memo_res{Memo: m})
	}
}

//Repositoryでidとowner_idを照会させるのでここでowner_idにリクエストを送ってきたuser_idを格納
func (handler *memoHandler) DeleteColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		//uriをバリデーション :memo_idにalphanum以外が入らないように
		uri_bid := &validators.MemoIDParams{}
		if err := c.ShouldBindUri(uri_bid); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		memo := models.ColorMemo{}
		user_id, _ := c.Get("user_id")
		memo.OwnerID, _ = user_id.(string)
		memo.ID = uri_bid.MemoID

		err := handler.memoController.DeleteColorMemo(&memo)
		//TODO: errにエラー定義したエラーを入れるようにする
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.JSON(http.StatusNoContent, "")
	}
}

//Repositoryでidとowner_idを照会させるのでここでowner_idにリクエストを送ってきたuser_idを格納
func (handler *memoHandler) EditColorMemo() gin.HandlerFunc {

	return func(c *gin.Context) {
		//body部分のバリデーション
		b_bid := &validators.ColorMemoValidator{}
		if err := c.Bind(b_bid); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		uri_bid := &validators.MemoIDParams{}
		if err := c.ShouldBindUri(uri_bid); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		user_id_, _ := c.Get("user_id")
		user_id, _ := user_id_.(string)
		var color_memo = &models.ColorMemo{
			ID:      uri_bid.MemoID,
			OwnerID: user_id,
			Color1:  b_bid.ColorMemo.Color1[1:],
			Color2:  b_bid.ColorMemo.Color2[1:],
		}

		m, err := handler.memoController.EditColorMemo(color_memo)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusForbidden, err.Error())
			return
		}
		c.JSON(http.StatusOK, memo_res{Memo: m})
	}
}
