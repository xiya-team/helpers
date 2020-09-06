package helpers

import "strings"

func ArrayChangeKeyase(arrayOrigin map[string]interface{}, changeTo int) map[string]interface{} {

	arrayReturn := make(map[string]interface{})
	for key, value := range arrayOrigin {

		keyNew := ""

		if changeTo == CASE_UPPER {
			keyNew = strings.ToUpper(key)
		} else if changeTo == CASE_LOWER {
			keyNew = strings.ToLower(key)
		} else {
			keyNew = key
		}

		arrayReturn[keyNew] = value.(interface{})
	}

	return arrayReturn
}
