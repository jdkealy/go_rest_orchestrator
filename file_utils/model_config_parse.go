package file_utils

import (
	jsLib "encoding/json"
	"fmt"
	"github.com/jdkealy/go_rails/types"
	"io/ioutil"
	"os"
)

func ParseJsonConfig(path string ) (*[]types.Fields, error) {
	var err error
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var fields []types.Fields

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	jsLib.Unmarshal(byteValue, &fields)
	return &fields, err
}