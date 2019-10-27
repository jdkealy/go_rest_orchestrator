package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenMainModel(s types.Schema){
	path := s.ModelPath + "/model.go"
	doGenAndSave(s, path, templates.ModelTypeTemplate)

	path = s.ModelPath + "/" + auto_migrate_file
	doGenAndSave(s, path, templates.AutoMigrateTemplate)
}