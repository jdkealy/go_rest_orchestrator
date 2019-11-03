package scaffold

import (
	"github.com/jdkealy/go_rails/file_utils"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"log"
)

func attachRouteToString(model string ) string {
	return "\trouter.Attach" + model + "Routes(c)\n"
}

func GenRouteFile(s types.Schema) {
	path := s.RoutesPath + "/" + s.PluralLowerModel + ".go"
	file_utils.DoGenAndSave(s, path, templates.RoutesTemplate)

	path = s.RoutesPath + "/" + gen_routes_path

	text := attachRouteToString(s.Model)

	err := file_utils.AppendStringToLine(path, gen_routes_substr, text)
	if err != nil {
		log.Fatal(err)
	}
	return
}