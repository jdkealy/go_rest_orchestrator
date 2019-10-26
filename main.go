package main

import (
	"github.com/jdkealy/go_rails/scaffold"
	"github.com/jdkealy/go_rails/types"
	"log"
	"os"
)



func cmd(s []string){
	switch s[1]{
	case "new": {
		initProject(s)
	}
	case "scaffold": {
		genScaffold(s)
	}
	default: {
		log.Println("Invalid command. Run help to see list of valid commands")
	}
	}
}

func initProject(s []string){
	d := types.InitPaths(s[2], s[3])
	os.MkdirAll(d.ProjectRoot, os.ModePerm)
	os.MkdirAll(d.ModelPath, os.ModePerm)
	os.MkdirAll(d.CmdPath, os.ModePerm)
	os.MkdirAll(d.RoutesPath, os.ModePerm)
	os.MkdirAll(d.TestsPath, os.ModePerm)
	os.MkdirAll(d.DbPath, os.ModePerm)
	os.MkdirAll(d.CachePath, os.ModePerm)
	scaffold.GenMainRoutes(d)
	scaffold.GenMainModel(d)
	scaffold.GenDb(d)
	scaffold.GenMain(d)

}

func genScaffold(s []string){
	schema := types.ParseSchema(s[2])
	scaffold.Routes(schema)
	scaffold.Models(schema)
}

func main() {
	//argsWithProg := os.Args
	argsWithProg := []string{"duh", "new", "jdkealy", "fuck_bfa" }
	cmd(argsWithProg)
}
