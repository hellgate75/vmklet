package vmkletio

import (
	"encoding/json"
	"encoding/xml"
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
		return err
	}
	return yaml.Unmarshal(data, out)
}

func GetJSONFromObj(m interface{}, prettify bool) []byte {
	var bytes []byte = make([]byte, 0)
	var err error
	if prettify {
		bytes, err = json.MarshalIndent(m, "", "  ")

	} else {
		bytes, err = json.Marshal(m)
	}
	if err != nil {
		return []byte{}
	}
	return bytes
}

func GetJSONFromElem(m interface{}, prettify bool) ([]byte, error) {
	if prettify {
		return json.MarshalIndent(m, "", "  ")
	}
	return json.Marshal(m)
}

func GetElemFromJSONBytes(data []byte, out *interface{}) error {
	return json.Unmarshal(data, out)
}

func GetElemFromJSONString(data string, out *interface{}) error {
	return json.Unmarshal([]byte(data), out)
}

func GetElemFromJSONFile(file string, out *interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

func GetXMLFromObj(m interface{}, prettify bool) []byte {
	var bytes []byte = make([]byte, 0)
	var err error
	if prettify {
		bytes, err = xml.MarshalIndent(m, "", "  ")

	} else {
		bytes, err = xml.Marshal(m)
	}
	if err != nil {
		return []byte{}
	}
	return bytes
}

func GetXMLFromElem(m interface{}, prettify bool) ([]byte, error) {
	if prettify {
		return xml.MarshalIndent(m, "", "  ")
	}
	return xml.Marshal(m)
}

func GetElemFromXMLBytes(data []byte, out *interface{}) error {
	return xml.Unmarshal(data, out)
}

func GetElemFromXMLString(data string, out *interface{}) error {
	return xml.Unmarshal([]byte(data), out)
}

func GetElemFromXMLFile(file string, out *interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, out)
}
