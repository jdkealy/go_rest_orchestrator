package main

import (
	"github.com/jdkealy/go_rest_orchestrator/file_utils"
	"log"
)

func main() {
	path := "/Users/johnkealy/go/src/github.com/jdkealy/bar/internal/pkg/models/auto_migrate_gen.go"
	substr := "AUTO_MIGRATE_GEN "
	file_utils.AppendStringToLine(path, substr, "TEST")
	log.Println("HI")
}
