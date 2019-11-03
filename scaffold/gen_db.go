
package scaffold

import (
	"github.com/jdkealy/go_rails/file_utils"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenDb(s types.Schema){
	path := s.DbPath + "/db.go"
	file_utils.DoGenAndSave(s, path, templates.DbTemplate)
}