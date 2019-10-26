package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"log"
	"text/template"
)

func GenMainModel(s types.Schema){
	t := genModelType(s)
	d1 := []byte(t.String())
	path := s.ModelPath + "/model.go"
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		log.Fatal("error writing main file")
	}
}

func genModelType(d types.Schema) bytes.Buffer {
	t := template.Must(template.New("models").Parse(templates.ModelTypeTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(buf.String())
	return buf
}