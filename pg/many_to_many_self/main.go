package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"
)

// Elem describes elements
type Elem struct {
	ID    int
	Elems []Elem `pg:"many2many:elem_to_elems,joinFK:sub_id"`
}

// ElemToElem describes element to element many-to-many-self relations.
type ElemToElem struct {
	ElemID int
	SubID  int
}

func init() {
	orm.RegisterTable((*ElemToElem)(nil))
}

func createManyToManySelfTables(db *pg.DB) error {
	models := []interface{}{
		(*Elem)(nil),
		(*ElemToElem)(nil),
	}
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     "demo",
		Password: "demo",
		Database: "demo",
		Addr:     "kube:5432",
	})
	defer db.Close()

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	if err := createManyToManySelfTables(db); err != nil {
		panic(err)
	}

	values := []interface{}{
		&Elem{ID: 1},
		&Elem{ID: 2},
		&Elem{ID: 3},
		&ElemToElem{ElemID: 1, SubID: 2},
		&ElemToElem{ElemID: 1, SubID: 3},
	}
	for _, v := range values {
		if err := db.Insert(v); err != nil {
			panic(err)
		}
	}

	elem := new(Elem)
	if err := db.Model(elem).Relation("Elems").First(); err != nil {
		panic(err)
	}

	fmt.Println("Elem", elem.ID)
	fmt.Println("Subelems", elem.Elems[0].ID, elem.Elems[1].ID)
}
