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
	schema := types.ParseSchema(s[2])
	scaffold.GenMain(schema)
}

func genScaffold(s []string){
	schema := types.ParseSchema(s[2])
	scaffold.Routes(schema)
	scaffold.Models(schema)
}

func main() {
	argsWithProg := os.Args
	cmd(argsWithProg)
}
