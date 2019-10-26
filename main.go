package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Type string
	Model string
	Owner string
	Name string
}

func main() {
	var d data
	flag.StringVar(&d.Owner, "owner", "owner", "Org name")
	flag.StringVar(&d.Project, "project", "project", "project name")
	flag.StringVar(&d.Model, "users", "User", "The subtype used for the queue being generated")
	flag.StringVar(&d.Name, "name", "FOO", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.Parse()

	t := template.Must(template.New("queue").Parse(queueTemplate))
	t.Execute(os.Stdout, d)
}

var queueTemplate = `
package routes

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/{{.Owner}}/{{.Project}}/internal/pkg/{{.Model}}"
	"log"
)

func (h *Routes) Attach{{.Model}}Routes(r *gin.Engine) {
	v1 := r.Group("/{{.Model}}")
	v1.GET("/", h.All{{.Model}})
	v1.GET("/new", h.New{{.Model}})
	v1.PUT("/:id", h.Edit{{.Model}})
	v1.DELETE("/:id", h.Delete{{.Model}})

}