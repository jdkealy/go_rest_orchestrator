package types

import (
	"flag"
	"github.com/gertd/go-pluralize"
	"strings"
)


type Schema struct {
	Type string
	Model string
	LowerModel string
	PluralModel string
	PluralLowerModel string
	Owner string
	Project string
	Name string
}

func ParseSchema(modelName string) Schema{
	var d Schema
	pluralize := pluralize.NewClient()
	pluralModelName := pluralize.Plural(modelName)
	lowercasePluralModelName := strings.ToLower(pluralModelName)
	d.PluralModel = pluralModelName
	d.LowerModel = strings.ToLower(modelName)
	d.PluralLowerModel = lowercasePluralModelName

	flag.StringVar(&d.Owner, "owner", "owner", "")
	flag.StringVar(&d.Project, "project", "project", "")
	flag.StringVar(&d.Model, "Model", modelName, "")
	flag.StringVar(&d.PluralModel, "PluralModel", d.PluralModel, "")
	flag.StringVar(&d.PluralLowerModel, "PluralLowerModel", d.PluralLowerModel, "")
	flag.StringVar(&d.Name, "name", "FOO", "")
	flag.Parse()

	return d
}