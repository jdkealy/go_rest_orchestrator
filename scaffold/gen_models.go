package scaffold

import (
	"fmt"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
	"os"
)

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

func GenModels(s types.Schema){
	path := s.ModelPath + "/" + s.PluralLowerModel + ".go"

	/* for now hardcode to test */
	doGenAndSave(s, path, templates.ModelTemplate)

	path = s.ModelPath + "/" + auto_migrate_file
	str := "\tdb.AutoMigrate(&" + s.Model + "{})\n"
	err := AppendStringToLine(path, gen_models_substr, str)
	if err != nil {
		//log.Fatal(err)
	}

}


func GenJs(s types.Schema){
	os.MkdirAll(s.JsModelsPath, os.ModePerm)
	os.MkdirAll(s.JsViewsPath, os.ModePerm)

	doGenAndSave(s, s.JsPageListPath, templates.JsListTemplate)
	doGenAndSave(s, s.JsPageNewPath, templates.JsLFormTemplate)
	doGenAndSave(s, s.JsModelPath, templates.JsModelTemplate)
	return
}