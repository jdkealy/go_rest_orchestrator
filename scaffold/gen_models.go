package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"log"
	"text/template"
)

func Models(s types.Schema){
	genModelTemplate(s)
}

func genModelTemplate(d types.Schema) bytes.Buffer {
	t := template.Must(template.New("models").Parse(templates.ModelTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(buf.String())
	return buf
}