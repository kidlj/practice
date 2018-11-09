package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"
)

// Profile describes user profile
type Profile struct {
	ID     int
	Lang   string
	Active bool
	UserID int
}

// User describes users
type User struct {
	ID       int
	Name     string
	Profiles []*Profile
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
		"CREATE TEMP TABLE users (id int, name text)",
		"CREATE TEMP TABLE profiles (id int, lang text, active bool, user_id int)",
		"INSERT INTO users VALUES (1, 'user 1'), (2, 'user 2')",
		"INSERT INTO profiles VALUES (1, 'en', TRUE, 1), (2, 'zh', TRUE, 1), (3, 'jp', FALSE, 1)",
	}

	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}

	var user User
	err := db.Model(&user).
		Column("user.*").
		Relation("Profiles", func(q *orm.Query) (*orm.Query, error) {
			return q.Where("active IS TRUE"), nil
		}).
		First()
	if err != nil {
		panic(err)
	}

	fmt.Println(user.ID, user.Name, user.Profiles[0], user.Profiles[1])
}
