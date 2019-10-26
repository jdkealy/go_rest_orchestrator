package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"log"
	"text/template"
)

func genRouteFile(d types.Schema) bytes.Buffer {
	t := template.Must(template.New("routes").Parse(templates.RoutesTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func Routes(s types.Schema){
	genRouteFile(s)
}