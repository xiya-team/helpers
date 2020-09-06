package helpers

func ArrayDiff(array1, array2 []interface{}) []interface{} {
	diffArr := []interface{}{}
	for _, val := range array1 {
		if InArray(val, array2) == false {
			diffArr = append(diffArr, val)
		}
	}
	return diffArr
}
