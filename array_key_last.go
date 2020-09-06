package helpers

import(
	// "reflect"
)
/**
 * array_key_first — Gets the first key of an array
 * @params map[interface{}]interface{} array An array.
 *
 * @return interface{}
 */
func ArrayKeyLast(arr map[interface{}]interface{}) interface{} {
	var k interface{}

	for key, _:= range arr {
		k = key
		break;
	}
	return k
}
