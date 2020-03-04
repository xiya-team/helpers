package helpers

// Date takes a PHP like date func to Go's time fomate
import (
	"time"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//dateTimeFormat 格式如："2006-01-02 15:04:05"
func Date(timestamp int64) string {
	return time.Unix(timestamp, 0).In(cstSh).Format(FORMAT_DATE)
}

//当前日期
func DateNow() string {
	return DateToday()
}

//当前日期
func DateToday() string {
	return time.Now().In(cstSh).Format(FORMAT_DATE)
}

//明天日期
func DateTomorrow() string {
	return time.Unix(Time() + 86400, 0).In(cstSh).Format(FORMAT_DATE)
}

//多少天后
func DateDaysafter(days int64, timestamp int64) string {
	return DateAfter(86400 * days, timestamp)
}

//多少天前
func DateDaysbefore(days int64, timestamp int64) string {
	return DateBefore(86400 * days, timestamp)
}

//多少小时后
func DateHoursafter(hours int64, timestamp int64) string {
	return DateAfter(3600 * hours, timestamp)
}

//多少小时前
func DateHoursbefore(hours int64, timestamp int64) string {
	return DateBefore(3600 * hours, timestamp)
}

//多少分钟后
func Date_minutesafter(minutes int64, timestamp int64) string {
	return DateAfter(60 * minutes, timestamp)
}

//多少分钟前
func DateMinutesbefore(minutes int64, timestamp int64) string {
	return DateBefore(60 * minutes, timestamp)
}

//多少秒后
func DateAfter(seconds int64, timestamp int64) string {
	return time.Unix(timestamp + seconds, 0).In(cstSh).Format(FORMAT_DATE)
}

//多少秒前
func DateBefore(seconds int64, timestamp int64) string {
	return time.Unix(timestamp - seconds, 0).In(cstSh).Format(FORMAT_DATE)
}