package main

import (
	"flag"
	"log"
	"os"
	"github.com/gertd/go-pluralize"
	"strings"
	"text/template"
)

type data struct {
	Type string
	Model string
	PluralModel string
	PluralLowerModel string
	Owner string
	Project string
	Name string
}

func cmd(s []string){
	switch s[1]{
	case "new": {
		log.Println("GENERATING NEW PROJECT")
	}
	case "scaffold": {
		scaffold(s)
	}
	default: {
		log.Println("Invalid command. Run help to see list of valid commands")
	}
	}
}

func scaffold(s []string){
	pluralize := pluralize.NewClient()
	modelName := s[2]
	pluralModelName := pluralize.Plural(modelName)
	lowercasePluralModelName := strings.ToLower(pluralModelName)

	var d data
	flag.StringVar(&d.Owner, "owner", "owner", "")
	flag.StringVar(&d.Project, "project", "project", "")
	flag.StringVar(&d.Model, "Model", modelName, "")
	flag.StringVar(&d.PluralModel, "PluralModel", pluralModelName, "")
	flag.StringVar(&d.PluralLowerModel, "PluralLowerModel", lowercasePluralModelName, "")
	flag.StringVar(&d.Name, "name", "FOO", "")
	flag.Parse()
	t := template.Must(template.New("routes").Parse(routesTemplate))
	t.Execute(os.Stdout, d)
}

func main() {
	argsWithProg := os.Args
	cmd(argsWithProg)
}

var routesTemplate = `
package routes

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/{{.Owner}}/{{.Project}}/internal/pkg/{{.Model}}"
	"log"
)

func (h *Routes) Attach{{.Model}}Routes(r *gin.Engine) {
	v1 := r.Group("/{{.PluralLowerModel}}")
	v1.GET("/", h.All{{.PluralModel}})
	v1.GET("/new", h.New{{.Model}})
	v1.PUT("/:id", h.Edit{{.Model}})
	v1.DELETE("/:id", h.Delete{{.Model}})
}`
