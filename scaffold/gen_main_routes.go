package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenMainRoutes(s types.Schema){
	path := s.RoutesPath + "/routes.go"
	doGenAndSave(s, path, templates.MainRoutes)

	path = s.RoutesPath + "/" + gen_routes_path
	doGenAndSave(s, path, templates.RegisterRoutesTemplate)
}