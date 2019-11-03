package scaffold

import (
	"fmt"
	"github.com/jdkealy/go_rails/file_utils"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"log"
	"os"
)

const js_import_substr =  "AUTOGEN_IMPORTS"
const js_route_config_substr =  "AUTOGEN_CONFIG"

func gorm(fields types.Fields) string {
	if fields.Gorm != "" {
		return fmt.Sprintf(`gorm:"%s"`, fields.Gorm)
	}
	return ""
}

func jsonStr(fields types.Fields) string {
	if fields.Json != "" {
		return fmt.Sprintf(`json:"%s"`, fields.Json)
	}
	return ""
}

func fieldToString(field types.Fields) string {
	return field.Name + " " + field.Type + "`" + gorm(field) + " " + jsonStr(field) + "`"
}

func fieldsToModel(fields []types.Fields) []string {
	var strs []string
	for _, item := range(fields){
		str := fieldToString(item)
		strs = append(strs, str)
	}
	return strs
}

func fieldsToRoute(fields []types.Fields) []string {
	var strs []string
	for _, item := range(fields){
		str := fieldToString(item)
		strs = append(strs, str)
	}
	return strs
}

func GenModels(s types.Schema){
	path := s.ModelPath + "/" + s.PluralLowerModel + ".go"

	/* for now hardcode to test */
	file_utils.DoGenAndSave(s, path, templates.ModelTemplate)

	path = s.ModelPath + "/" + auto_migrate_file
	str := "\tdb.AutoMigrate(&" + s.Model + "{})\n"
	err := file_utils.AppendStringToLine(path, gen_models_substr, str)
	if err != nil {
		//log.Fatal(err)
	}
}


func GenJs(s types.Schema){
	os.MkdirAll(s.JsModelsPath, os.ModePerm)
	os.MkdirAll(s.JsViewsPath, os.ModePerm)
	file_utils.DoGenAndSave(s, s.JsPageListPath, templates.JsListTemplate)
	file_utils.DoGenAndSave(s, s.JsPageFieldsPath, templates.JsFieldsTemplate)
	file_utils.DoGenAndSave(s, s.JsPageNewPath, templates.JsLFormTemplate)
	file_utils.DoGenAndSave(s, s.JsModelPath, templates.JsModelTemplate)
	str := fmt.Sprintf(`{route: "/%s", cmp: %s, label: '%s' },
	`, s.PluralLowerModel, s.Model, s.PluralModel)
	err := file_utils.AppendStringToLine(s.JsRoutesConfigPath, js_route_config_substr, str )
	if err != nil {
		log.Fatal(err)
	}

	file_utils.CleanJs(s.JsRoutesConfigPath)

	// generate the string with compnent here
	str = fmt.Sprintf(
		`import %s from '../pages/%s/list'
	`, s.Model, s.Model)
	err = file_utils.AppendStringToLine(s.JsRoutesConfigPath, js_import_substr, str)
	if err != nil {
		log.Fatal(err)
	}

	return
}
