package path

import (
	"fmt"
	"net/url"
	"path"
	"testing"
)

func Test_join(t *testing.T) {
	p := path.Join("", "a")
	if p != "a" {
		t.Error("failed")
	}
}

func Example_filepath() {
	p := "/logs%2Flogs.489072"
	r, err := url.PathUnescape(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	// Output:
	// /logs/logs.489072
}
