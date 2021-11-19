package main

import "testing"

// You can index a nil map
func Test_nil(t *testing.T) {
	var m map[string]string
	if m["a"] != "" {
		t.Error("failed")
	}
}
