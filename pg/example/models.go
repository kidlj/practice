package example

import (
	"fmt"
	"time"
)

type Genre struct {
	ID        int
	Name      string
	Rating    int    `sq:"-"`
	Books     []Book `pg:"many2many:book_genres"`
	ParentID  int
	Subgenres []Genre `pg:"fk:parent_id"`
}

type Book struct {
	ID        int
	Title     string
	AuthorID  int
	Author    *Author
	EditorID  int
	Editor    *Author
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time `sql:"default:current_timestamp"`

	Genres       []Genre `pg:"many2many:book_genres"`
	Translations []Translation
	Comments     []Comment `pg:"polymorphic:trackable_"`
}

// func (b *Book) BeforeInsert(db orm.DB) error {
// 	if b.CreatedAt.IsZero() {
// 		b.CreatedAt = time.Now()
// 	}
// 	return nil
// }

type Author struct {
	ID       int
	Name     string
	Books    []*Book // has many relation
	AvatarID int
	Avatar   Image // has one relation
}

type Image struct {
	ID   int
	Path string
}

func (a Author) String() string {
	return fmt.Sprintf("Author<ID=%d Name=%q", a.ID, a.Name)
}

type BookGenre struct {
	BookID       int `sql:",pk"`
	Book         *Book
	GenreID      int `sql:",pk"`
	Genre        *Genre
	Genre_Rating int // belongs to and is copied to Genre model
}

type Translation struct {
	ID       int
	BookID   int       `sql:"unique:book_id_lang"`
	Book     *Book     // has one relation
	Lang     string    `sql:"unique:book_id_lang"`
	Comments []Comment `pg:",polymorphic:trackable_"`
}

type Comment struct {
	TrackableID   int
	TrackableType string
	Text          string
}
