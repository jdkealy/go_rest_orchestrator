package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"log"
)

func GenRouteFile(s types.Schema) {
	path := s.RoutesPath + "/" + s.PluralLowerModel + ".go"
	doGenAndSave(s, path, templates.RoutesTemplate)

	path = s.RoutesPath + "/" + gen_routes_path
	text := "\trouter.Attach" + s.Model + "Routes(c)\n"
	err := AppendStringToLine(path, gen_routes_substr, text)
	if err != nil {
		log.Fatal(err)
	}
	return
}