package json

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Example_configureCertificates() {
	s := configureCertificates()
	fmt.Println(s)
	// Output: xx
}

func Example_Pointer() {
	type T struct {
		I *int   `json:"i,omitempty"`
		S string `json:"s,omitempty"`
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

func Example_Map() {
	m := map[string]struct{}{
		"a": {},
		"b": {},
	}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%q\n", b)
	// Output:
	// "{\"a\":{},\"b\":{}}"
}

// omitempty has no effect on unmarshalling.
func Test_omitempty(t *testing.T) {
	type S struct {
		Payload struct {
			Op    string `json:"op"`
			After struct {
				ID int `json:"id"`
			} `json:"after"`
		} `json:"payload"`
	}

	var s = new(S)

	data := `{ "payload":{"op": "d"}}`
	err := json.Unmarshal([]byte(data), s)
	if err != nil {
		t.Error("unmarshall failed")
	}
	if s.Payload.Op != "d" {
		t.Error("test failed")
	}
}
