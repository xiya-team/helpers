package helpers

// ArrayMerge array_merge()
func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

//对于切片不建议使用本方法，而要直接使用 append
func ArrayMergeString(arr1, arr2 []string) []string {
	return append(arr1, arr2...)
}

func MapMerge(arr1, arr2 map[string]interface{}) map[string]interface{} {
	newArr := arr1
	for key, value := range arr2 {
		newArr[key] = value
	}

	return newArr
}

func MapMergeString(arr1, arr2 map[string]string) map[string]string {
	newArr := arr1
	for key, value := range arr2 {
		newArr[key] = value
	}

	return newArr
}

func MapMergeKeyInterface(arr1, arr2 map[interface{}]interface{}) map[interface{}]interface{} {
	newArr := arr1
	for key, value := range arr2 {
		newArr[key] = value
	}

	return newArr
}
