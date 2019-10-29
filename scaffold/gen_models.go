package scaffold

import (
	"fmt"
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

type fields struct {
	Name string
	Type string
	Gorm string
	Json string
}

func gorm(fields fields) string {
	if fields.Gorm != "" {
		return fmt.Sprintf(`gorm:"%s"`, fields.Gorm)
	}
	return ""
}

func json(fields fields) string {
	if fields.Json != "" {
		return fmt.Sprintf(`json:"%s"`, fields.Json)
	}
	return ""
}

func fieldToString(field fields) string {
	return field.Name + " " + field.Type + "`" + gorm(field) + " " + json(field) + "`"
}

func fieldsToModel(fields []fields) []string {
	var strs []string
	for _, item := range(fields){
		str := fieldToString(item)
		strs = append(strs, str)
	}
	return strs
}

func GenModels(s types.Schema){
	path := s.ModelPath + "/" + s.PluralLowerModel + ".go"
	doGenAndSave(s, path, templates.ModelTemplate)

	path = s.ModelPath + "/" + auto_migrate_file
	str := "\tdb.AutoMigrate(&" + s.Model + "{})\n"
	err := AppendStringToLine(path, gen_models_substr, str)
	if err != nil {
		//log.Fatal(err)
	}

}