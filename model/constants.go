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
	"json",
	"xml",
	"yaml",
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
	for i := 0; i < len(FormatTypeList); i++ {
		if FormatTypeList[i] == *Type {
			return i
		}
	}
	return -1
}

func (String *FormatString) Parse() FormatType {
	var index int = String.Index()
	if index >= 0 {
		return FormatTypeList[index]
	}
	return NoFormat
}

func (Type *FormatString) Index() int {
	for i := 0; i < len(FormatTypeStringList); i++ {
		if FormatTypeStringList[i] == *Type {
			return i
		}
	}
	return -1
}

type StreamType int
type StreamString string

const (
	SSHStream StreamType = iota
	NoStream
)

var StreamTypeList []StreamType = []StreamType{
	SSHStream,
	NoStream,
}

var StreamTypeStringList []StreamString = []StreamString{
	"SSH",
	"",
}

func (Type *StreamType) String() string {
	var index int = Type.Index()
	if index >= 0 && index < len(StreamTypeStringList) {
		return string(StreamTypeStringList[index])
	}
	return ""
}

func (Type *StreamType) Index() int {
	for i := 0; i < len(StreamTypeList); i++ {
		if StreamTypeList[i] == *Type {
			return i
		}
	}
	return -1
}

func (String *StreamString) Parse() StreamType {
	var index int = String.Index()
	if index >= 0 {
		return StreamTypeList[index]
	}
	return NoStream
}

func (Type *StreamString) Index() int {
	for i := 0; i < len(StreamTypeStringList); i++ {
		if StreamTypeStringList[i] == *Type {
			return i
		}
	}
	return -1
}

type CommandType int
type CommandString string

const (
	ShellCommand CommandType = iota
	NoCommand
)
var CommandTypeList []CommandType = []CommandType{
	ShellCommand,
	NoCommand,
}


var CommandTypeStringList []CommandString = []CommandString{
	"SHELL",
	"",
}

func (Type *CommandType) String() string {
	var index int = Type.Index()
	if index >= 0 && index < len(CommandTypeStringList) {
		return string(CommandTypeStringList[index])
	}
	return ""
}

func (Type *CommandType) Index() int {
	for i := 0; i < len(CommandTypeList); i++ {
		if CommandTypeList[i] == *Type {
			return i
		}
	}
	return -1
}

func (String *CommandString) Parse() CommandType {
	var index int = String.Index()
	if index >= 0 {
		return CommandTypeList[index]
	}
	return NoStream
}

func (Type *CommandString) Index() int {
	for i := 0; i < len(CommandTypeStringList); i++ {
		if CommandTypeStringList[i] == *Type {
			return i
		}
	}
	return -1
}
