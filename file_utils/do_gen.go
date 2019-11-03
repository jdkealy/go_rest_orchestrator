package file_utils

import (
	"bytes"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"log"
	"text/template"
)

func DoGenAndSave(s types.Schema, path string , templateName string ){
	t := DoGenTemplate(s, templateName )
	d1 := []byte(t.String())
	err := ioutil.WriteFile(path, d1, 0644)
	log.Println("GENERATING", path)
	if err != nil {
		log.Fatal(err)
	}
}

func DoGenTemplate(d types.Schema, templateName string ) bytes.Buffer {
	t := template.Must(template.New("template").Parse(templateName))
	var buf bytes.Buffer
	err := t.Execute(&buf, d)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}