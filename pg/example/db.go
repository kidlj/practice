package example

import (
	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     "demo",
		Password: "demo",
		Database: "demo",
	})
}
