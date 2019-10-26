package types

import (
	"flag"
	"github.com/gertd/go-pluralize"
	"os"
	"strings"
)


type Schema struct {
	FullFilePath string
	DependencyPath string
	CmdPath string
	ModelPath string
	RoutesPath string
	TestsPath string
	DbPath string
	CachePath string
	GoRoot string
	GitPath string
	ProjectRoot string
	GitRoutesPath string
	GitDbPath string
	GitModelsPath string

	Type string
	Model string
	LowerModel string
	PluralModel string
	PluralLowerModel string
	Owner string
	Project string
	Name string
}

func initPaths(owner string, project string, d Schema ) Schema {
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

func InitPaths(owner string, project string )Schema {
	var d Schema
	return initPaths(owner, project, d)
}

func ParseSchema(modelName string) Schema{
	d := InitPaths("jdkealy", "bar")
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