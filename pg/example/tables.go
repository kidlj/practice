package example

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func createSchema(db *pg.DB, models []interface{}) error {
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func dropSchema(db *pg.DB, models []interface{}) error {
	for _, model := range models {
		err := db.DropTable(model, &orm.DropTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
