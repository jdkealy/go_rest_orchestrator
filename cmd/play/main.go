package main

import (
	"github.com/jdkealy/go_rails/scaffold"
	"log"
)

func main() {
	path := "/Users/johnkealy/go/src/github.com/jdkealy/bar/internal/pkg/models/auto_migrate_gen.go"
	substr := "AUTO_MIGRATE_GEN "
	scaffold.AppendStringToLine(path, substr, "TEST")
	log.Println("HI")
}