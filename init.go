package helpers

import "time"

var (
	cstSh           *time.Location
	FORMAT_DATE     string = "2006-01-02"
	FORMAT_DATETIME string = "2006-01-02 15:04:05"
	CASE_UPPER int = 1	//转为大写
	CASE_LOWER int = 0	//转为小写
	RET_STRING = 0		//结果为STRING
	RET_INT = 1			//结果为INT
	RET_INT64 = 2		//结果为INT64
)