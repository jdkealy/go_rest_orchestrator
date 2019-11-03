package scaffold

import (
	"github.com/jdkealy/go_rest_orchestrator/file_utils"
	"github.com/jdkealy/go_rest_orchestrator/templates"
	"github.com/jdkealy/go_rest_orchestrator/types"
)

func GenMainModel(s types.Schema){
	path := s.ModelPath + "/js_model.go"
	file_utils.DoGenAndSave(s, path, templates.ModelTypeTemplate)

	path = s.ModelPath + "/" + auto_migrate_file
	file_utils.DoGenAndSave(s, path, templates.AutoMigrateTemplate)
}
