package scaffold

import (
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
	root := go_root + "/src/" + git_path
	cmd := root + "/cmd"
	models := root + "/internal/pkg/models"
	routes := root + "/routes"
	git_routes_path := git_path + "/routes"
	git_db_path := git_path + "/db"
	git_models_path := git_path + "/internal/pkg/models"
	tests := root + "/tests"


	db := root + "/db"
	cache := root + "/cache"

	d.ProjectRoot = root
	d.GitPath = git_path
	d.GitRoutesPath = git_routes_path
	d.GitModelsPath = git_models_path
	d.CachePath = cache
	d.DbPath = db
	d.GitDbPath = git_db_path
	d.TestsPath  = tests
	d.RoutesPath = routes
	d.ModelPath = models
	d.CmdPath = cmd

	return d
}

func InitPaths(owner string, project string ) types.Schema {
	var d types.Schema
	return initPaths(owner, project, d)
}

func ParseSchema(owner string, project string, modelName string, path string) types.Schema{
	var modelDef *[]types.Fields
	structFields, err := parseJsonConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	fields := fieldsToModel(*structFields)

	modelDef, err = parseJsonConfig(path)
	log.Println(modelDef, err)
	d := InitPaths(owner, project)
	pluralize := pluralize.NewClient()
	pluralModelName := pluralize.Plural(modelName)
	lowercasePluralModelName := strings.ToLower(pluralModelName)
	d.StructFields = fields
	d.PluralModel = pluralModelName
	d.LowerModel = strings.ToLower(modelName)
	d.PluralLowerModel = lowercasePluralModelName

	views_path := fmt.Sprintf(`%s/views`, d.ProjectRoot)
	js_model_path := fmt.Sprintf(`%s/src/models/%s.js`, views_path, modelName )
	js_new_page_path := fmt.Sprintf(`%s/src/%s/new.js`, views_path, modelName )
	js_list_page_path := fmt.Sprintf(`%s/src/%s/list.js`, views_path, modelName )
	js_models_path := fmt.Sprintf(`%s/src/models/`, views_path )
	js_views_path :=  fmt.Sprintf(`%s/src/%s/`, views_path )

	d.JsViewsPath = js_views_path
	d.JsModelsPath = js_models_path
	d.JsModelPath = js_model_path
	d.JsPageListPath = js_list_page_path
	d.JsPageNewPath = js_new_page_path
	flag.StringVar(&d.Owner, "owner", "owner", "")
	flag.StringVar(&d.Project, "project", "project", "")
	flag.StringVar(&d.Model, "Model", modelName, "")
	flag.StringVar(&d.PluralModel, "PluralModel", d.PluralModel, "")
	flag.StringVar(&d.PluralLowerModel, "PluralLowerModel", d.PluralLowerModel, "")
	flag.StringVar(&d.Name, "name", "FOO", "")
	flag.Parse()
	return d
}