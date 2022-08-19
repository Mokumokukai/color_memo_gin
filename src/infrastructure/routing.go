package infrastructure

import (
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure/middleware"
	"github.com/Mokumokukai/color_memo_gin/src/registry"

	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *DB
	Port string
	Gin  *gin.Engine
}

func NewRouting(db *DB, reg registry.IInteractor) *Routing {
	c := NewConfig()

	r := &Routing{
		DB:   db,
		Port: c.Routing.Port,
		Gin:  gin.Default(),
	}
	r.setRouting(reg)
	return r

}

func (r *Routing) setRouting(reg registry.IInteractor) {
	uh := reg.NewUserHandler()
	mh := reg.NewColorMemoHandler()
	th := reg.NewTagHandler()
	r.Gin.GET("/memos", mh.GetColorMemos())
	r.Gin.GET("/tags", th.GetTags())
	r.Gin.GET("/users", uh.GetUsers())

	jwtm := middleware.NewJWTMiddlwareHandler("/go/src/firebase-adminsdk.json", "color-memo-auth", r.DB.Connection)
	r.Gin.POST("/auth/signup", jwtm.SetUID(), uh.Register())

	r.Gin.Use(jwtm.SetUserID())
	r.Gin.POST("/memos", mh.CreateColorMemo())
	r.Gin.POST("/memos/:memo_id/duplicate", mh.DuplicateColorMemo())
	r.Gin.POST("/memos/:memo_id/delete", mh.DeleteColorMemo())
	r.Gin.POST("/memos/:memo_id/edit", mh.EditColorMemo())

	r.Gin.POST("/tags", th.CreateTag())

}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}

func (r *Routing) Stop() {
}
