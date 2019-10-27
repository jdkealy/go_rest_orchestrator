package scaffold

import (
	"github.com/jdkealy/go_rails/templates"
	"github.com/jdkealy/go_rails/types"
)

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