package helpers

import "fmt"

// ArrayColumn array_column()
// ArrayColumn — Return the values from a single column in the input array
func ArrayColumn(arrayMap map[string]map[string]interface{}, columnKey string) (r []interface{}) {

	for _, input := range arrayMap {
		if v, ok := input[columnKey]; ok {
			r = append(r, v)
		}
	}

	return
}

func ArrayColumnRetArrayString(array []map[string]interface{}, column string) []string {

	retArray := []string{}

	for _, value := range array {
		retArray = append(retArray, fmt.Sprintf("%v", value[column]))
	}

	return retArray
}

func ArrayColumnMapStringStringRetArrayString(array []map[string]string, column string) []string {

	retArray := []string{}

	for _, value := range array {
		retArray = append(retArray, value[column])
	}

	return retArray
}
