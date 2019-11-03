package scaffold

import (
	"github.com/jdkealy/go_rails/file_utils"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenMain(s types.Schema){
	path := s.ProjectRoot + "/main.go"
	file_utils.DoGenAndSave(s, path, templates.MainTemplate)
}