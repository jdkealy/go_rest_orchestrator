package scaffold

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/jdkealy/go_rails/types"
	"log"
	"os"
	"strings"
)

func initPaths(owner string, project string, d types.Schema ) types.Schema {
	go_root := os.Getenv("GOPATH")
	git_path := "github.com/" + owner + "/"  + project
	project_root := go_root + "/src/" + git_path
	cache := project_root + "/cache"
	d.ProjectRoot = project_root
	d.ViewPath = fmt.Sprintf(`%s/views`, d.ProjectRoot)
	d.JsModelsPath = fmt.Sprintf(`%s/src/models/`, d.ViewPath )
	d.GitPath = git_path
	d.GitRoutesPath = git_path + "/routes"
	d.GitModelsPath =  git_path + "/internal/pkg/models"
	d.CachePath = cache
	d.DbPath = project_root + "/db"
	d.GitDbPath = git_path + "/db"
	d.TestsPath  = fmt.Sprintf(`%s/tests`, project_root)
	d.RoutesPath = project_root + "/routes"
	d.ModelPath = project_root + "/internal/pkg/models"
	d.CmdPath = project_root + "/cmd"

	return d
}

func InitPaths(owner string, project string ) types.Schema {
	var d types.Schema
	return initPaths(owner, project, d)
}

func makeRoutesJson(m []types.Fields ) []types.RoutesJsonField{
	var routesJsonFields []types.RoutesJsonField
	for _, i := range m{
		var rjs = types.RoutesJsonField{
			Name: i.Name,
			Id:   i.Name,
			Cmp:  fmt.Sprintf(i.Name),
		}
		routesJsonFields = append(routesJsonFields, rjs)
	}
	return routesJsonFields
}

func fieldsToJsonFields(m []types.Fields ) []types.JsModelsJsonFields{
	var routesJsonFields []types.JsModelsJsonFields
	for _, i := range m{
		var rjs = types.JsModelsJsonFields{
			Name: i.Name,
			Id:   i.Name,
			Label:  i.Name,
		}
		routesJsonFields = append(routesJsonFields, rjs)
	}
	return routesJsonFields
}

func ParseSchema(owner string, project string, modelName string, path string) types.Schema{
	d := InitPaths(owner, project)
	structFields, err := parseJsonConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	fields := fieldsToModel(*structFields)
	var modelDef *[]types.Fields
	modelDef, err = parseJsonConfig(path)
	log.Println(modelDef)

	jsonFields := fieldsToJsonFields(*structFields)
	f, _ := json.Marshal(jsonFields)
	d.JsFieldsConfig = string(f)

	jsonRoutesFields := makeRoutesJson(*structFields)
	f, _ = json.Marshal(jsonRoutesFields)
	d.JsRoutesConfig = string(f)

	pluralize := pluralize.NewClient()
	pluralModelName := pluralize.Plural(modelName)

	d.StructFields = fields
	d.PluralModel = pluralModelName
	d.LowerModel = strings.ToLower(modelName)
	d.PluralLowerModel = strings.ToLower(pluralModelName)

	d.JsViewsPath = fmt.Sprintf(`%s/src/pages/%s/`, d.ViewPath, modelName )
	d.JsPageFieldsPath = fmt.Sprintf(`%s/src/pages/%s/fields.js`, d.ViewPath, modelName )
	d.JsRoutesConfigPath = fmt.Sprintf(`%s/src/components/routes_config.js`, d.ViewPath )
	d.JsModelPath = fmt.Sprintf(`%s/src/models/%s.js`, d.ViewPath, modelName )
	d.JsPageListPath = fmt.Sprintf(`%s/src/pages/%s/list.js`, d.ViewPath, modelName )
	d.JsPageNewPath = fmt.Sprintf(`%s/src/pages/%s/new.js`, d.ViewPath, modelName )
	flag.StringVar(&d.Owner, "owner", "owner", "")
	flag.StringVar(&d.Project, "project", "project", "")
	flag.StringVar(&d.Model, "Model", modelName, "")
	flag.StringVar(&d.PluralModel, "PluralModel", d.PluralModel, "")
	flag.StringVar(&d.PluralLowerModel, "PluralLowerModel", d.PluralLowerModel, "")
	flag.StringVar(&d.Name, "name", "FOO", "")
	flag.Parse()
	return d
}