package file_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../")
	userPath = "test_files/models/user.json"
)

type ModelConfigParseSuite struct {
	suite.Suite
}

func (suite *ModelConfigParseSuite) TestGoFmt() {
	s := fmt.Sprintf(`%s%s%s`, "FOO", "BAR", "MEOW")
	assert.Equal(suite.T(), "FOOBARMEOW", s)

}
func (suite *ModelConfigParseSuite) TestModels() {
	path := fmt.Sprintf(`%s/%s`, Root, userPath)
	fields, err := ParseJsonConfig(path)
	log.Println(fields)
	assert.Equal(suite.T(), err, nil)
}
func TestModelConfigParseSuite(t *testing.T) {
	suite.Run(t, new(ModelConfigParseSuite))
}