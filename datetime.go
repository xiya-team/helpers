package helpers

import (
	"time"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//dateTimeFormat 格式如："2006-01-02 15:04:05"
func Datetime(timestamp int64) string {
	return time.Unix(timestamp, 0).In(cstSh).Format(FORMAT_DATETIME)
}

//当前日期时间
func DatetimeNow() string {
	return DatetimeToday()
}

//当前日期时间
func DatetimeToday() string {
	return time.Now().In(cstSh).Format(FORMAT_DATETIME)
}

//明天时间
func DatetimeTomorrow() string {
	return time.Unix(Time()+86400, 0).In(cstSh).Format(FORMAT_DATETIME)
}

//多少天后
func DatetimeDaysafter(days int64, timestamp int64) string {
	return DatetimeAfter(86400*days, timestamp)
}

//多少天前
func DatetimeDaysbefore(days int64, timestamp int64) string {
	return DatetimeBefore(86400*days, timestamp)
}

//多少小时后
func DatetimeHoursafter(hours int64, timestamp int64) string {
	return DatetimeAfter(3600*hours, timestamp)
}

//多少小时前
func DatetimeHoursbefore(hours int64, timestamp int64) string {
	return DatetimeBefore(3600*hours, timestamp)
}

//多少分钟后
func DatetimeMinutesafter(minutes int64, timestamp int64) string {
	return DatetimeAfter(60*minutes, timestamp)
}

//多少分钟前
func DatetimeMinutesbefore(minutes int64, timestamp int64) string {
	return DatetimeBefore(60*minutes, timestamp)
}

//多少秒后
func DatetimeAfter(seconds int64, timestamp int64) string {
	return time.Unix(timestamp+seconds, 0).In(cstSh).Format(FORMAT_DATETIME)
}

//多少秒前
func DatetimeBefore(seconds int64, timestamp int64) string {
	return time.Unix(timestamp-seconds, 0).In(cstSh).Format(FORMAT_DATETIME)
}
