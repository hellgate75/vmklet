package test

import (
	"github.com/hellgate75/vmklet/contrib"
	"github.com/hellgate75/vmklet/model"
	"github.com/hellgate75/vmklet/test/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewStateParserJSON(t *testing.T) {
	var author utils.Person = utils.Person{
		Name:    "Fabrizio",
		Surname: "Torelli",
		Age:     42,
	}
	parser, err := contrib.NewParser(model.JSONFormat)
	parser.Prettify = true
	parser.Structure = author
	assert.Nil(t, err, "New JSON Parser must return nil error")
	file, err := ioutil.TempFile("", "golang_test_parser")
	defer func(file string) {
		os.Remove(file)
	}(file.Name())
	parser.Save(file.Name())
	_, err = os.Stat(file.Name())
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
	expectedFileContent := "{\n  \"Name\": \"Fabrizio\",\n  \"Surname\": \"Torelli\",\n  \"Age\": 42\n}"
	assert.Equal(t, expectedFileContent, string(bytes), "Parsed File must be JSON must be same")
}

func TestNewStateParserYAML(t *testing.T) {
	var author utils.Person = utils.Person{
		Name:    "Fabrizio",
		Surname: "Torelli",
		Age:     42,
	}
	parser, err := contrib.NewParser(model.YAMLFormat)
	parser.Prettify = true
	parser.Structure = author
	assert.Nil(t, err, "New JSON Parser must return nil error")
	file, err := ioutil.TempFile("", "golang_test_parser")
	defer func(file string) {
		os.Remove(file)
	}(file.Name())
	parser.Save(file.Name())
	_, err = os.Stat(file.Name())
	assert.Nil(t, err, "New YAML Parsed file must exists")
	var authorParsed utils.Person
	parser.Structure = authorParsed
	bytes, err := ioutil.ReadFile(file.Name())
	err = parser.Load(file.Name())
	assert.Nil(t, err, "New YAML Parser should be able to parse file")
	authorX := parser.Structure.(map[interface{}]interface{})
	authorName := authorX["Name"]
	authorSurname := authorX["Surname"]
	authorAge := authorX["Age"]
	assert.Equal(t, author.Name, authorName, "Parsed Name must be same")
	assert.Equal(t, author.Surname, authorSurname, "Parsed Surname must be same")
	assert.Equal(t, int(author.Age), authorAge, "Parsed Age must be same")
	expectedFileContent := "Name: Fabrizio\nSurname: Torelli\nAge: 42\n"
	assert.Equal(t, expectedFileContent, string(bytes), "Parsed File must be JSON must be same")
}

func TestNewStateParserXML(t *testing.T) {
	var author utils.Person = utils.Person{
		Name:    "Fabrizio",
		Surname: "Torelli",
		Age:     42,
	}
	parser, err := contrib.NewParser(model.XMLFormat)
	parser.Prettify = false
	parser.Structure = author
	assert.Nil(t, err, "New XML Parser must return nil error")
	file, err := ioutil.TempFile("", "golang_test_parser")
	defer func(file string) {
		os.Remove(file)
	}(file.Name())
	parser.Save(file.Name())
	_, err = os.Stat(file.Name())
	assert.Nil(t, err, "New JSON Parsed file must exists")
	var authorParsed utils.Person
	parser.Structure = authorParsed
	bytes, err := ioutil.ReadFile(file.Name())
	// Specialized function to parse (unmarshall) the Person struct
	err = utils.GetPersonFromXMLFile(file.Name(), &authorParsed)
	assert.Nil(t, err, "New XML Parser should be able to parse file")
	authorName := authorParsed.Name
	authorSurname := authorParsed.Surname
	authorAge := authorParsed.Age
	assert.Equal(t, author.Name, authorName, "Parsed Name must be same")
	assert.Equal(t, author.Surname, authorSurname, "Parsed Surname must be same")
	assert.Equal(t, author.Age, authorAge, "Parsed Age must be same")
	expectedFileContent := "<Person><Name>Fabrizio</Name><Surname>Torelli</Surname><Age>42</Age></Person>"
	assert.Equal(t, expectedFileContent, string(bytes), "Parsed File must be JSON must be same")
}
