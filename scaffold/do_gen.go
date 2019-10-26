package scaffold

import (
	"bytes"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"log"
	"text/template"
)

func doGenAndSave(s types.Schema, templateName string, path string ){
	t := doGenTemplate(s, templateName )
	d1 := []byte(t.String())
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		log.Fatal("error writing main file")
	}
}

func doGenTemplate(d types.Schema, templateName string ) bytes.Buffer {
	t := template.Must(template.New("template").Parse(templateName))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}