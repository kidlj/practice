package path

import (
	"path"
	"testing"
)

func Test_join(t *testing.T) {
	p := path.Join("", "a")
	if p != "a" {
		t.Error("failed")
	}
}
