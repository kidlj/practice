package template

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func Blocks() {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}}{{join . ", "}}{{"\n"}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.ExecuteTemplate(os.Stdout, "list", guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.ExecuteTemplate(os.Stdout, "master", guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.ExecuteTemplate(os.Stdout, "list", guardians); err != nil {
		log.Fatal(err)
	}
}
