package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"log"
	"text/template"
)

func GenMain(s types.Schema){
	t := genMain(s)
	d1 := []byte(t.String())
	path := s.ProjectRoot + "/main.go"
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		log.Fatal("error writing main file")
	}
}

func genMain(d types.Schema) bytes.Buffer {
	t := template.Must(template.New("models").Parse(templates.MainTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}