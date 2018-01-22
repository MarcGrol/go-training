package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name      string
	Interests []string
	Children  []Child
}

type Child struct {
	Name string
	Age  int
}

func (c *Child) NameAndAge() string {
	return fmt.Sprintf("%s (%d)", c.Name, c.Age)
}

var p = Person{
	Name:      "Marc Grol",
	Interests: []string{"Running", "Golang"},
	Children: []Child{
		{Name: "Pien", Age: 5},
		{Name: "Tijl", Age: 12},
		{Name: "Freek", Age: 18},
	},
}

func main() {
	const tpl = `
<html>
	<body>
		<p>Hi, I'm {{.Name}}.
		<p>I like {{range .Interests}}{{.}}, {{end}}
		<ul>
		{{range .Children}}<li>{{.NameAndAge}}</li>{{end}}
		</ul>
	</body>
</html>`

	t, _ := template.New("person").Parse(tpl)
	_ = t.Execute(os.Stdout, p)
}
