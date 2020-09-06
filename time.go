package helpers

import "time"

// Time time()
func Time() int64 {
	return time.Now().Unix()
}
