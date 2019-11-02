package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

func GenJs(s types.Schema){
	os.MkdirAll(s.JsModelsPath, os.ModePerm)
	os.MkdirAll(s.JsViewsPath, os.ModePerm)
	doGenAndSave(s, s.JsPageListPath, templates.JsListTemplate)
	doGenAndSave(s, s.JsPageNewPath, templates.JsLFormTemplate)
	doGenAndSave(s, s.JsModelPath, templates.JsModelTemplate)
	return
}