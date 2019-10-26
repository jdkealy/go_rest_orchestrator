package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenMain(s types.Schema){
	path := s.ProjectRoot + "/main.go"
	doGenAndSave(s, path, templates.MainTemplate)
}