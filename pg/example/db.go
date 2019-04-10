package example

import (
	"fmt"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func init() {
	DB = pg.Connect(&pg.Options{
		User:     "demo",
		Password: "demo",
		Database: "demo",
	})
	fmt.Println(DB)
}
