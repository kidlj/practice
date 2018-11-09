package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
)

// Profile describes user profile
type Profile struct {
	ID     int
	Lang   string
	UserID int
	User   *User
}

// User describes users
type User struct {
	ID   int
	Name string
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
		"CREATE TEMP TABLE profiles (id int, lang text, user_id int)",
		"INSERT INTO users VALUES (1, 'user 1'), (2, 'user 2')",
		"INSERT INTO profiles VALUES (1, 'en', 1), (2, 'zh', 2)",
	}

	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}

	var profiles []Profile
	err := db.Model(&profiles).
		Column("profile.*").
		Relation("User").
		Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(profiles), "results")
	fmt.Println(profiles[0].ID, profiles[0].Lang, profiles[0].User)
	fmt.Println(profiles[1].ID, profiles[1].Lang, profiles[1].User)
}
