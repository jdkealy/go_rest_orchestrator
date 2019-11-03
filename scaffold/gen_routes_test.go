package scaffold

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
}

func (suite *ExampleTestSuite) TestModels() {
	assert.Equal(suite.T(), 5, 4)
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
