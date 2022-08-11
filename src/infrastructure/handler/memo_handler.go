package handler

import (
	"github.com/Mokumokukai/color_memo_gin/src/adaptor/controllers"
	"github.com/gin-gonic/gin"
)

type memoHandler struct {
	memoController controllers.IColorMemoController
}

type IColorMemoHandler interface {
	GetColorMemos() gin.HandlerFunc
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
		c.JSON(200, m)
	}
}
