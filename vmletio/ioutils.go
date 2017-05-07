package vmletio

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetYAMLFromObj(m interface{}) []byte {
	bytes, err := yaml.Marshal(m)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func GetYAMLFromElem(m interface{}) ([]byte, error) {
	return yaml.Marshal(m)
}

func GetElemFromYAMLBytes(data []byte, out *interface{}) error {
	return yaml.Unmarshal(data, out)
}

func GetElemFromYAMLString(data string, out *interface{}) error {
	return yaml.Unmarshal([]byte(data), out)
}

func GetElemFromYAMLFile(file string, out *interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return  err
	}
	return yaml.Unmarshal(data, out)
}
