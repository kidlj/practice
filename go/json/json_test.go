package json

import (
	"encoding/json"
	"fmt"
	"log"
)

func Example_configureCertificates() {
	s := configureCertificates()
	fmt.Println(s)
	// Output: xx
}

func Example_Pointer() {
	type T struct {
		I *int   `json:"i,omitempty"`
		S string `json:"s,omitemtpy"`
	}
	j := `
	{
		"s": "hello"
	}`
	var t = new(T)
	err := json.Unmarshal([]byte(j), t)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(t.S, t.I)
	bytes, err := json.Marshal(t)
	fmt.Println(string(bytes))
	// Output:
	// {}
}
