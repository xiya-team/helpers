package helpers

// ArrayPush - Push one or more elements onto the end of array
func ArrayPush(s *[]interface{}, elements ...interface{}) int {
	*s = append(*s, elements...)
	return len(*s)
}