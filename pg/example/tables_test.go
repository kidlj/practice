package example

import "fmt"

func Example_createSchema() {
	models := []interface{}{(*Book)(nil), (*Author)(nil), (*Genre)(nil), (*BookGenre)(nil), (*Translation)(nil), (*Comment)(nil), (*Image)(nil)}
	err := createSchema(DB, models)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
	// Output:
	// ok
}

func Example_dropSchema() {
	models := []interface{}{(*Book)(nil), (*Author)(nil), (*Genre)(nil), (*BookGenre)(nil), (*Translation)(nil), (*Comment)(nil), (*Image)(nil)}
	err := dropSchema(DB, models)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
	// Output:
	// ok
}
