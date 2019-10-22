// https://www.v2ex.com/t/598241
package main

import "fmt"

type Plugin interface {
	Pr()
}

var Plugins = make(map[string]Creator)

func Add(name string, creator Creator) {
	Plugins[name] = creator
}

type Test struct {
	Name string
}

func (t Test) Pr() {
	fmt.Println(t.Name)
}

type Creator func() Plugin

func CreatePlugin(name string) Plugin {
	creator := Plugins[name]
	plugin := creator()
	return plugin
}

func init() {
	Add("test", func() Plugin {
		return &Test{"mellon"}
	})
}

func main() {
	pluginName := "test"
	plugin := CreatePlugin(pluginName)
	plugin.Pr()
}
