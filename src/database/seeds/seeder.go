package main

import (
	"log"

	"github.com/Mokumokukai/color_memo_gin/src/database/seeds/seeds"
	"github.com/Mokumokukai/color_memo_gin/src/infrastructure"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := infrastructure.NewDB()

	for _, seed := range seeds.All() {
		if err := seed.Run(db.Connection); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
