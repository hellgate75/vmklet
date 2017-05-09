package utils

import (
	"io/ioutil"
	"encoding/xml"
)

type Person struct {
	Name    string `json:"Name" xml:"Name" yaml:"Name"`
	Surname string `json:"Surname" xml:"Surname" yaml:"Surname"`
	Age     int    `json:"Age" xml:"Age" yaml:"Age"`
}
func GetPersonFromXMLFile(file string, out *Person) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, out)
}
