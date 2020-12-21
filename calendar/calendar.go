package calendar

import (
	"fmt"
	"math"
	"time"
)

// Second to String
func SecondToTime(t int) (res string) {
	var day, hour, minute int
	if t >= 86400 {
		day = int(math.Floor(float64(t / 86400)))
		t = (t % 86400)
	}

	if t >= 3600 {
		hour = int(math.Floor(float64(t / 3600)))
		t = (t % 3600)
	}

	if t >= 60 {
		minute = int(math.Floor(float64(t / 60)))
		t = (t % 60)
	}

	res = fmt.Sprintf("%d天%d小时%d分钟%d秒", day, hour, minute, t)

	return
}

// Format By Timestamp
func FormatByTimestamp(t string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	times, _ := time.ParseInLocation(FMT.ToString(), t, loc)
	return times.Unix()
}

// Timestamp By Time
func TimestampByFormat(t int64, f TimeFormat) string {
	unix := time.Unix(t, 0)
	return unix.Format(f.ToString())
}

// Format Time
func GetTime(f TimeFormat) string {
	return t().Format(f.ToString())
}

// Timestamp
func Timestamp() int64 {
	return t().Unix()
}

// Get Year
func GetYear() int {
	return t().Year()
}

// Get Month
func GetMonth() string {
	month := t().Month().String()
	return month
}

// Get Day
func GetDay() int {
	return t().Day()
}

// Get Hour
func GetHour() int {
	return t().Hour()
}

// Get Minute
func GetMinute() int {
	return t().Minute()
}

// Time
func t() time.Time {
	return time.Now()
}
