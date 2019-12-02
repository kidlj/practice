package template

import (
	"os"
	"text/template"
)

type Numbers struct {
	X *int
	Y *int
}

// This example demos that 0 valued *int pointer is not 'empty'.
func Example_interface() {
	var i = 3
	var j = 0
	var n = &Numbers{
		&i, &j,
	}
	tmpl, err := template.New("test").Parse("{{if .X}}{{.X}}{{end}}, {{if .Y}}{{.Y}}{{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, n)
	if err != nil {
		panic(err)
	}
	// Output:
	// 3, 0
}
