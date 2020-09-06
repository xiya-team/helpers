package helpers

import(
	// "reflect"
)

func ArraySearchString(needle string, haystack []string) int {
	length := len(haystack)
	 
	for i:= 0;i< length;i++ {
		if (haystack[i] == needle) {
			return i
		}
	}

	return -1
}

func ArraySearch(needle interface{}, haystack []interface{}) int {

	length := len(haystack)

	for i:= 0; i<length; i++ {
		if (haystack[i] == needle) {
			return i
		}
	}

	return -1
}
