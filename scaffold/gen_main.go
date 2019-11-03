package scaffold

import (
	"github.com/jdkealy/go_rest_orchestrator/file_utils"
	"github.com/jdkealy/go_rest_orchestrator/templates"
	"github.com/jdkealy/go_rest_orchestrator/types"
)

func GenMain(s types.Schema){
	path := s.ProjectRoot + "/main.go"
	file_utils.DoGenAndSave(s, path, templates.MainTemplate)
}
