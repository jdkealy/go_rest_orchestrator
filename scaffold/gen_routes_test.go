package scaffold

import (
	"github.com/gertd/go-pluralize"
	"github.com/jdkealy/go_rails/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
}

func (suite *ExampleTestSuite)TestTemplate(t *testing.T) {
	var d types.Schema
	pluralize := pluralize.NewClient()
	modelName := "Antelope"
	pluralModelName := pluralize.Plural(modelName)
	lowercasePluralModelName := strings.ToLower(pluralModelName)
	d.PluralModel = pluralModelName
	d.PluralLowerModel = lowercasePluralModelName
	template := genRouteFile(d)
	templateToString := template.String()
	assert.True(suite.T(), strings.Contains(templateToString, "Antelopes"), "should contain antelopes")
}

func (suite *ExampleTestSuite)TestGenModel(t *testing.T) {
	d := types.ParseSchema("Kitten")
	template := genRouteFile(d)
	templateToString := template.String()
	assert.True(suite.T(), strings.Contains(templateToString, "Antelopes"), "should contain kittens")

}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}