package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"log"
	"text/template"
)

func GenMainRoutes(s types.Schema){
	t := genMainRoutes(s)
	d1 := []byte(t.String())
	path := s.RoutesPath + "/routes.go"
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		log.Fatal("error writing main file")
	}
}

func genMainRoutes(d types.Schema) bytes.Buffer {
	t := template.Must(template.New("routes").Parse(templates.MainRoutes))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}