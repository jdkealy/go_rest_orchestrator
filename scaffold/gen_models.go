package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func Models(s types.Schema){
	path := s.ModelPath + s.PluralLowerModel + ".go"
	doGenAndSave(s, path, templates.ModelTemplate)
}