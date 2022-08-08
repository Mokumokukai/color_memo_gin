package main

import (
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure"
	"github.com/Mokumokukai/color_memo_gin/src/registry"
)

func main() {
	db := infrastructure.NewDB()
	defer infrastructure.CloseDB(db)
	reg := registry.NewInteractor(db.Connection)
	r := infrastructure.NewRouting(db, reg)
	r.Run()

}
