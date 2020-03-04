package helpers

import "sort"

// ArrayKeys array_keys()
func ArrayKeys(elements map[interface{}]interface{}) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

//key的类型为int
func ArrayKeysInt(array map[int]interface{}) []int {
	retArray := []int{}
	for key,_ := range array {
		retArray = append(retArray, key)
	}
	sort.Ints(retArray)

	return retArray
}

//key的类型为interface{}
func ArrayKeysInterface(array map[interface{}]interface{}) []interface{} {
	retArray := []interface{}{}
	for key,_ := range array {
		retArray = append(retArray, key)
	}

	return retArray
}