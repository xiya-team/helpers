package helpers

func ArrayKeyExistsString(needle string, haystack map[string]interface{}) bool {
	var isExist = false
	_, isExist = haystack[needle]
	return isExist
}
