package libs

import (
	_"fmt"
	"time"
	."project/libs/lib_conv"
)

func Time(param ...interface{}) interface{} {
	var ret interface{}

	switch(len(param)) {
	default:
		ret = time.Now().Unix()
	case 0:
		ret = time.Now().Unix()
	case 1:
		switch(param[0]) {
		default:
			ret = time.Now().Unix()
		case "unix":
			ret = time.Now().Unix()
		case "now":
			ret = time.Now()
		case "utc":
			ret = time.Unix(time.Now().Unix(), 0)
		case "h", "hr", "hh", "hour", "hours", "時", "小時", "时", "小时":
			ret = time.Now().Hour()
		case "m", "min", "minute", "minutes", "分", "分鐘", "分钟":
			ret = time.Now().Minute()
		case "s", "sec", "second", "seconds", "秒":
			ret = time.Now().Second()
		case "y", "yr", "year", "years", "年":
			ret = time.Now().Year()
		case "M", "mo", "mos", "mon", "mons", "month", "months", "月":
			ret = int(time.Now().Month())
		case "monthName":
			ret = time.Now().Month()
		case "月份", "月分":
			chtMon := func(t int) string {
				chtMon := []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
				return chtMon[int(time.Now().Month())-1]
			}

			ret = chtMon(int(time.Now().Month()))
		case "d", "day", "日":
			ret = time.Now().Day()
		case "when":
			now := Time("hour").(int)
			if (now >= 0 && now <=12) {
				ret = "AM"
			} else {
				ret = "PM"
			}

			return ret
		case "上下午", "上午", "下午":
			now := Time("hour").(int)
			if (now >= 0 && now <=12) {
				ret = "上午"
			} else {
				ret = "下午"
			}

			return ret
		case "time":
			now := Time("hour").(int)
			if (now > 0 && now < 6) {
				ret = "Evening"
			} else if (now >= 6 && now <= 11) {
				ret = "Morning"
			} else if (now == 12) {
				ret = "Noon"
			} else if (now >= 13 && now < 19) {
				ret = "Afternoon"
			} else if (now >= 19 && now < 23) {
				ret = "Night"
			} else {
				ret = "Midnight"
			}

			return ret
		case "時份", "时份":
			now := Time("hour").(int)
			if (now >= 1 && now < 5) {
				ret = "凌晨"
			} else if (now >= 5 && now < 7) {
				ret = "清晨"
			} else if (now >= 7 && now < 9) {
				ret = "早上"
			} else if (now >= 9 && now < 12) {
				ret = "上午"
			} else if (now == 12) {
				ret = "中午"
			} else if (now >= 13 && now < 18) {
				ret = "下午"
			} else if (now >= 18 && now < 20) {
				ret = "傍晚"
			} else if (now >= 20 && now < 23) {
				ret = "晚上"
			} else if (now == 23) {
				ret = "深夜"
			} else {
				ret = "午夜"
			}

			return ret
		}
	case 2:
		interval := ConvInf2Int(param[1])
		if chk := TypeOfInt(interval); chk {
			ret = 0
		}
		switch(param[0]) {
		default:
			ret = time.Now().Unix()
		case "format":
			ret = time.Now().Format(ConvInf2Str(param[1]))
		case "beforeYear":
			ret = time.Now().AddDate(-interval, 0, 0).Format("2006-01-02")
		case "afterYear":
			ret = time.Now().AddDate(interval, 0, 0).Format("2006-01-02")
		case "beforeMonth":
			ret = time.Now().AddDate(0, -interval, 0).Format("2006-01-02")
		case "afterMonth":
			ret = time.Now().AddDate(0, interval, 0).Format("2006-01-02")
		case "beforeDay":
			ret = time.Now().AddDate(0, 0, -interval).Format("2006-01-02")
		case "afterDay":
			ret = time.Now().AddDate(0, 0, interval).Format("2006-01-02")
		case "beforeHour":
			ret = time.Now().Add(-time.Duration(interval) * time.Hour).Format("15:04:05")
		case "afterHour":
			ret = time.Now().Add(time.Duration(interval) * time.Hour).Format("15:04:05")
		case "beforeMin":
			ret = time.Now().Add(-time.Duration(interval) * time.Minute).Format("15:04:05")
		case "afterMin":
			ret = time.Now().Add(time.Duration(interval) * time.Minute).Format("15:04:05")
		case "beforeSec":
			ret = time.Now().Add(-time.Duration(interval) * time.Second).Format("15:04:05")
		case "afterSec":
			ret = time.Now().Add(time.Duration(interval) * time.Second).Format("15:04:05")
		}
	}

	return ret
}

