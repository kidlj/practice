package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
)

// Item describes item
type Item struct {
	ID       int
	Items    []Item `pg:"fk:parent_id"`
	ParentID int
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

	qs := []string{
		"CREATE TEMP TABLE items (id int, parent_id int)",
		"INSERT INTO items VALUES (1, NULL), (2, 1), (3, 1)",
	}

	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}

	var item Item
	err := db.Model(&item).
		Column("item.*").
		Relation("Items").
		First()
	if err != nil {
		panic(err)
	}

	fmt.Println("Item", item.ID)
	fmt.Println("Subitems", item.Items[0].ID, item.Items[1].ID)
}
