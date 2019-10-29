package scaffold

import (
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type TestGenModelsSuite struct {
	suite.Suite
}

func (suite *TestGenModelsSuite) TestModels() {
	 field := fields{
		Name: "foo",
		Gorm: "BAR",
		Json: "meow",
	}
	var fields []fields
	fields = append(fields, field)
	model := fieldsToModel(fields)
	log.Println(model)
}
func TestGenModels(t *testing.T) {
	suite.Run(t, new(TestGenModelsSuite))
}