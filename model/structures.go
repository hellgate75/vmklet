package model

import (
	"errors"
	"io/ioutil"
	"vmklet/vmkletio"
)

type Parser struct {
	Format    FormatType
	Structure interface{}
	Prettify 	bool
}

func (Parser *Parser) Save(file string) error {
	if Parser.Format == NoFormat {
		return errors.New("Unknown Parser!!")
	}
	var bytes []byte = make([]byte, 0)
	var err error = errors.New("Unknown Parser!!")
	if Parser.Format == JSONFormat {
		bytes, err = vmkletio.GetJSONFromElem(Parser.Structure, Parser.Prettify)
	} else if Parser.Format == XMLFormat {
		bytes, err = vmkletio.GetYAMLFromElem(Parser.Structure)
	} else if Parser.Format == YAMLFormat {
		bytes, err = vmkletio.GetXMLFromElem(Parser.Structure, Parser.Prettify)
	}
	if err == nil {
		err = ioutil.WriteFile(file, bytes, 0666)
	}
	return err
}

func (Parser *Parser) Load(file string) error {
	if Parser.Format == NoFormat {
		return errors.New("Unknown Parser!!")
	}
	var bytes []byte = make([]byte, 0)
	var err error = errors.New("Unknown Parser!!")
	if Parser.Format == JSONFormat {
		err = vmkletio.GetElemFromJSONFile(file, &Parser.Structure)
	} else if Parser.Format == XMLFormat {
		err = vmkletio.GetElemFromYAMLFile(file, &Parser.Structure)
	} else if Parser.Format == YAMLFormat {
		err = vmkletio.GetElemFromXMLFile(file, &Parser.Structure)
	}
	if err == nil {
		err = ioutil.WriteFile(file, bytes, 0666)
	}
	return err
}