package scaffold

import (
	"github.com/jdkealy/go_rails/file_utils"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenMainModel(s types.Schema){
	path := s.ModelPath + "/js_model.go"
	file_utils.DoGenAndSave(s, path, templates.ModelTypeTemplate)

	path = s.ModelPath + "/" + auto_migrate_file
	file_utils.DoGenAndSave(s, path, templates.AutoMigrateTemplate)
}