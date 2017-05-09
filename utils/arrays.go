package utils

func IndexOf(list []interface{}, elem interface{}) int {
	for i := 0; i < len(list); i++ {
		if list[i] == elem {
			return i
		}
	}
	return -1
}
