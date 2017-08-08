package util

import "time"

/*
CSTOffset 北京时间相对于UTC时间偏移量
*/
const (
	CSTOffset int64 = HourSec * 8
	DaySec    int64 = HourSec * 24
	WeekSec   int64 = DaySec * 7
	HourSec   int64 = 3600
)

// TimeStamp2CST 时间戳转为北京时间
func TimeStamp2CST(t int64) time.Time {
	temp := time.Unix(t+CSTOffset, 0)
	return temp.In(time.UTC)
}

// CST2TimeStamp 北京时间转为时间戳
func CST2TimeStamp(CSTTime time.Time) int64 {
	return CSTTime.Unix() - CSTOffset
}

// GetTodayZero 获取当天零点时间
func GetTodayZero(ts int64, zeroSec int64) int64 {
	// zeroSec%DaySec 保证传入的zeroSec必须在24小时内
	return ts - ((ts+CSTOffset)%DaySec-zeroSec%DaySec+DaySec)%DaySec
}

// IsToday 是否当天
func IsToday(t int64, zeroTime int64) bool {
	diff := t - zeroTime
	return diff >= 0 && diff < int64(24*time.Hour.Seconds())
}

// IsCurrentMonth 是否当月
func IsCurrentMonth(last, now, zeroSec int64) bool {
	lastZeroTimeStamp := GetTodayZero(last, zeroSec)
	lastCSTTime := TimeStamp2CST(lastZeroTimeStamp)

	nowZeroTimeStamp := GetTodayZero(now, zeroSec)
	nowCSTTime := TimeStamp2CST(nowZeroTimeStamp)

	return lastCSTTime.Year() == nowCSTTime.Year() && lastCSTTime.Month() == nowCSTTime.Month()
}

// IsCurrentWeek 是否本周
func IsCurrentWeek(last, now, zeroSec int64) bool {
	return GetMondayZeroTimeStamp(last, zeroSec) == GetMondayZeroTimeStamp(now, zeroSec)
}

// GetMondayZeroTimeStamp 获取每周星期一时间零点时间戳
func GetMondayZeroTimeStamp(ts, zeroSec int64) int64 {
	// 获取当前时间零点
	zeroTimeStamp := GetTodayZero(ts, zeroSec)
	zeroTime := TimeStamp2CST(zeroTimeStamp)
	zeroDay := (int64(time.Monday) + zeroSec/DaySec) % 7
	// 计算差值
	dif := (int64(zeroTime.Weekday()) - zeroDay + 7) % 7

	return zeroTimeStamp - dif*int64(time.Hour.Seconds())*24
}

// GetTimeStampFromCNTime 获取中国上海市区时间的时间戳
func GetTimeStampFromCNTime(strTime string) int64 {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.UTC)

	return CST2TimeStamp(t)
}

// GetCNStrTimeFromTimeStamp 获取时间戳的上海时间
func GetCNStrTimeFromTimeStamp(timeStamp int64) string {
	t := TimeStamp2CST(timeStamp)
	return t.Format("2006-01-02 15:04:05")
}
