package scaffold

import (
	"github.com/jdkealy/go_rest_orchestrator/file_utils"
	"github.com/jdkealy/go_rest_orchestrator/templates"
	"github.com/jdkealy/go_rest_orchestrator/types"
)

func GenMainRoutes(s types.Schema){
	path := s.RoutesPath + "/routes.go"
	file_utils.DoGenAndSave(s, path, templates.MainRoutes)

	path = s.RoutesPath + "/" + gen_routes_path
	file_utils.DoGenAndSave(s, path, templates.RegisterRoutesTemplate)
}
