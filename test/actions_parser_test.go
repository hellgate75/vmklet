package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vmklet/actions"
	"vmklet/test/utils"
	"vmklet/model"
	"io/ioutil"
	"os"
)

func TestNewStateParserJSON(t *testing.T) {
	var author utils.Person = utils.Person{
		Name:    "Fabrizio",
		Surname: "Torelli",
		Age: 			42,
	}
	parser, err := actions.NewParser(model.JSONFormat)
	parser.Prettify = true
	parser.Structure = author
	assert.Nil(t, err, "New JSON Parser must return nil error")
	file,err := ioutil.TempFile("", "golang_test_parser")
	defer func(file string) {
		os.Remove(file)
	}(file.Name())
	parser.Save(file.Name())
	_,err = os.Stat(file.Name())
	assert.Nil(t, err, "New JSON Parsed file must exists")
	var authorParsed utils.Person
	parser.Structure = authorParsed
	bytes, err := ioutil.ReadFile(file.Name())
	err = parser.Load(file.Name())
	assert.Nil(t, err, "New JSON Parser should be able to parse file")
	authorX := parser.Structure.(map[string]interface{})
	authorName := authorX["Name"]
	authorSurname := authorX["Surname"]
	authorAge := authorX["Age"]
	assert.Equal(t, author.Name, authorName, "Parsed Name must be same")
	assert.Equal(t, author.Surname, authorSurname, "Parsed Surname must be same")
	assert.Equal(t, float64(author.Age), authorAge, "Parsed Age must be same")
	expectedFileContent:="{\n  \"Name\": \"Fabrizio\",\n  \"Surname\": \"Torelli\",\n  \"Age\": 42\n}"
	assert.Equal(t, expectedFileContent, string(bytes), "Parsed File must be JSON must be same")
}
