package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"log"
	"text/template"
)

func GenDb(s types.Schema){
	t := genDb(s)
	d1 := []byte(t.String())
	path := s.DbPath + "/db.go"
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		log.Fatal("error writing main file")
	}
}

func genDb(d types.Schema) bytes.Buffer {
	t := template.Must(template.New("models").Parse(templates.DbTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}