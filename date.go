package helpers

// Date takes a PHP like date func to Go's time fomate
import (
	"time"
)

// Date date()
// Date("02/01/2006 15:04:05 PM", 1524799394)
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}
