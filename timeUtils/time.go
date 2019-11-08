package timeUtils

import (
	"time"
)

//Unix时间戳(秒)
func NowUnix() int64 {
	return time.Now().Unix()
}

//Unix时间戳(毫秒)
func NowUnixMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//Unix时间戳(微秒)
func NowUnixMicro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

//Unix时间戳(纳秒)
func NowUnixNano() int64 {
	return time.Now().UnixNano()
}
