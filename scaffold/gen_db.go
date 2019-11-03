
package scaffold

import (
	"github.com/jdkealy/go_rest_orchestrator/file_utils"
	"github.com/jdkealy/go_rest_orchestrator/templates"
	"github.com/jdkealy/go_rest_orchestrator/types"
)

func GenDb(s types.Schema){
	path := s.DbPath + "/db.go"
	file_utils.DoGenAndSave(s, path, templates.DbTemplate)
}
