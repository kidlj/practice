package stdnet

import (
	"fmt"
	"net/url"
	"testing"
)

func Example_Parse() {
	u, err := url.Parse("kidlj.com/wiki/python")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", u)
	// Output:
	// kidlj.com
}

func Test_Parse(t *testing.T) {
	u, err := url.Parse("kidlj.com/wiki/python")
	if err != nil {
		t.Error("parse url error")
	}
	if u.Host != "kidlj.com" {
		fmt.Print(u.Hostname())
		t.Error("parse url host error")
	}
}
