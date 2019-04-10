package example

import "fmt"

func ExampleInsertBook() {
	book := &Book{
		Title:    "The Go Programming Language",
		AuthorID: 1,
		EditorID: 2,
	}
	err := DB.Insert(book)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
	// Output:
	// ok
}

func ExampleInsertAuthor() {
	author1 := &Author{
		Name:     "Jian Li",
		AvatarID: 1,
	}
	author2 := &Author{
		Name:     "Mellon Collie",
		AvatarID: 2,
	}
	err := DB.Insert(author1, author2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
	// Output:
	// ok
}
