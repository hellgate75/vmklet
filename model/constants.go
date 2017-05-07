package model

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

func (Type *FormatType) String() string {
	var index int = Type.Index()
	if index >= 0 && index < len(FormatTypeStringList) {
		return string(FormatTypeStringList[index])
	}
	return ""
}

func (Type *FormatType) Index() int {
	for i:=0; i<len(FormatTypeList); i++ {
		if FormatTypeList[i] == *Type {
			return i
		}
	}
	return  -1
}

func (String *FormatString) Parse() FormatType {
	var index int = String.Index()
	if index >= 0 {
		return FormatTypeList[index]
	}
	return NoFormat
}

func (Type *FormatString) Index() int {
	for i:=0; i<len(FormatTypeStringList); i++ {
		if FormatTypeStringList[i] == *Type {
			return i
		}
	}
	return  -1
}
