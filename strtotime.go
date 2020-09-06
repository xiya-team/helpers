package helpers

import (
	"fmt"
	"time"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//例如：2019-01-31 00:00:00
func StrToTime(strtime interface{}) int64 {
	return StrToTimeDateTime(strtime)
}

//例如：2019-01-31 00:00:00
func StrToTimeDateTime(strtime interface{}) int64 {
	return strToTimeDateTime(strtime)
}

//例如：2019-01-31
func StrToTimeDate(strtime interface{}) int64 {
	return strToTimeDateTime(fmt.Sprintf("%v 00:00:00", strtime))
}

//外部不可以直接使用， 例如：2019-01-31 00:00:00
func strToTimeDateTime(strtimeI interface{}) int64 {
	strtime := fmt.Sprintf("%v", strtimeI)
	time1, err := time.ParseInLocation(FORMAT_DATETIME, strtime, cstSh)
	if err != nil {

		return 0
	}
	return time1.Unix()
}
