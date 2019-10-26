package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenDb(s types.Schema){
	path := s.DbPath + "/db.go"
	doGenAndSave(s, path, templates.DbTemplate)
}