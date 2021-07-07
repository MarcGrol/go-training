package main

import (
	"fmt"
	"html/template"
	"log"
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
	// START OMIT
	const tpl = `<html><body>
		<p>Hi, I'm {{.FullName}}.
		I like {{range $i, $el := .Interests}}{{if $i}} and {{end}}{{$el}}{{end}}</p>
		<ul>
		{{range .Children -}}
			<li>{{.NameAndAge}}</li>
		{{end -}}
		</ul>
	</body></html>`

	t, err := template.New("person").Parse(tpl)
	if err != nil { // Templates are not type strong! Need unit tests to prove ...
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, p)
	if err != nil { // Needs verification!!!
		// END OMIT
		log.Fatal(err)
	}

}
