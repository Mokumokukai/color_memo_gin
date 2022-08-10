package infrastructure

import (
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
	r.Gin.GET("/users", uh.GetUsers())
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}

func (r *Routing) Stop() {
}
