package utils

import (
	"strings"
	"strconv"
	"syscall"
)

func StrPad(instr string, capping int) string {
	strLen := len(instr)
	if strLen == capping {
		return instr
	} else {
		if strLen < capping {
			padding := strings.Repeat(" ", (capping - strLen))
			return instr + padding
		} else {
			val := instr[0:(capping - 2)]
			val += ".."
			return val
		}
	}
}

func CorrectInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(CorrectInput(s))
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func ReverseString(str string) string {
	bytesArray, err := syscall.ByteSliceFromString(str)
	if err != nil {
		return str
	}
	size := len(bytesArray)
	cycle := len(bytesArray) / 2
	for i := 0; i < cycle; i++ {
		var b byte = bytesArray[i]
		bytesArray[i] = bytesArray[size-1-i]
		bytesArray[size-1-i] = b
	}
	return string(bytesArray)
}


func GetBoolean(input string) bool {
	return CorrectInput(input) == "true"
}

func BoolToString(input bool) string {
	if input {
		return "yes"
	} else {
		return "no"
	}
}

