package template

import (
	"log"
	"os"
	"text/template"
)

type Conf struct {
	Servers []Server
}

type Server struct {
	Name     string
	TLS      TLSConf
	Location []Location
}

type TLSConf struct {
	Cert string
	Key  string
}

type Location struct {
	Path     string
	Canary   bool
	Services []Service
}

type Service struct {
	Name string
	Port int
}

func Example_trim() {
	conf := Conf{
		Servers: []Server{
			{
				Name: "www.test.com",
				TLS: TLSConf{
					Cert: "test-cert",
					Key:  "test-key",
				},
				Location: []Location{
					{
						Path:   "/header",
						Canary: true,
						Services: []Service{
							{
								Name: "echo-demo-v1",
								Port: 80,
							},
							{
								Name: "echo-demo-v2",
								Port: 80,
							},
						},
					},
					{
						Path:   "/cookie",
						Canary: true,
						Services: []Service{
							{
								Name: "echo-demo-v3",
								Port: 80,
							},
							{
								Name: "echo-demo-v4",
								Port: 80,
							},
						},
					},
				},
			},
		},
	}

	tmpl, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		log.Fatal("parse template failed", err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "conf", conf)
	// err = tmpl.Execute(os.Stdout, conf)
	if err != nil {
		log.Fatal("execute template failed", err)
	}
	// Output:
	// debug
}
