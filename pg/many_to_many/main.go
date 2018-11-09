package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/go-pg/pg"
)

// Subject describes subjects
type Subject struct {
	ID       int
	Title    string
	Articles []Article `pg:"many2many:article_to_subjects"`
}

// Article describes articles
type Article struct {
	ID       int
	Title    string
	Subjects []*Subject `pg:"many2many:article_to_subjects"`
}

// ArticleToSubject describe Article and Subject many-to-many relations
type ArticleToSubject struct {
	ArticleID int
	SubjectID int
}

func init() {
	orm.RegisterTable((*ArticleToSubject)(nil))
}

func createManyToManyTables(db *pg.DB) error {
	models := []interface{}{
		(*Article)(nil),
		(*Subject)(nil),
		(*ArticleToSubject)(nil),
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

	if err := createManyToManyTables(db); err != nil {
		panic(err)
	}

	values := []interface{}{
		&Article{ID: 1, Title: "hello"},
		&Article{ID: 2, Title: "world"},
		&Subject{ID: 1, Title: "4G"},
		&Subject{ID: 2, Title: "5G"},
		&ArticleToSubject{ArticleID: 1, SubjectID: 1},
		&ArticleToSubject{ArticleID: 1, SubjectID: 2},
		&ArticleToSubject{ArticleID: 2, SubjectID: 1},
		&ArticleToSubject{ArticleID: 2, SubjectID: 2},
	}
	for _, v := range values {
		if err := db.Insert(v); err != nil {
			panic(err)
		}
	}

	article := new(Article)
	if err := db.Model(article).Relation("Subjects").First(); err != nil {
		panic(err)
	}

	fmt.Println("Article", article.ID, article.Title, "Subjects", article.Subjects[0].ID, article.Subjects[1].ID)

	subject := new(Subject)
	if err := db.Model(subject).Relation("Articles").First(); err != nil {
		panic(err)
	}
	fmt.Println("Subject", subject.ID, subject.Title, "Articles", subject.Articles[0].ID, subject.Articles[1].ID)
}
