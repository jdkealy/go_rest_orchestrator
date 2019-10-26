package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenRouteFile(s types.Schema) {
	path := s.RoutesPath + "/" + s.PluralLowerModel + ".go"
	doGenAndSave(s, path, templates.RoutesTemplate)
	return
}
