package model

import "strings"

type FormatType int
type FormatString string

const (
	JSONFormat FormatType = iota
	XMLFormat
	YAMLFormat
	NoFormat
)

var FormatTypeList []FormatType = []FormatType{
	JSONFormat,
	XMLFormat,
	YAMLFormat,
	NoFormat,
}

var FormatTypeStringList []FormatString = []FormatString{
	"JSON",
	"XML",
	"YAML",
	"",
}

func (Type *FormatType) String() string{
	if int(Type) < len(FormatTypeStringList) {
		return  string(FormatTypeStringList[int(Type)])
	}
	return ""
}

func (String *FormatString) Parse() FormatType{
	var index int = strings.Index(FormatTypeList, String)
	if index >= 0 {
		return  FormatTypeList[index]
	}
	return NoFormat
}
