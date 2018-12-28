package utils

import (
	"fmt"
	"time"
)

type Test struct {
	S string
}

func main() {
	timeLayout := "2006-01-02 15:04:05"
	start := time.Now().Unix()
	start = 1544671203
	fmt.Println(time.Unix(start, 0).Format(timeLayout))
	//time.Sleep(time.Duration(5)*time.Second)
	endtime := time.Now().Unix()
	d, h, m, s := GetTime(endtime - start)
	fmt.Printf("%d天%d小时%d分钟%d秒", d, h, m, s)
}

func GetTime(sinter int64) (d, h, m, s int64) {
	if day, interval := getBasic(sinter, 86400); day > 0 {
		d = day
		h, m = oneDay(interval)
	} else {
		h, m = oneDay(interval)
	}
	s = sinter - (d * 86400) - (h * 3600) - (m * 60)
	return
}

func oneDay(interval int64) (h, m int64) { // 一天内的小时分钟
	h = interval / 3600
	if h == 0 {
		m = interval / 60
		return
	} else {
		interval = interval - (3600 * h)
		m = interval / 60
		return
	}
	return
}

func getBasic(sinter, tt int64) (num, interval int64) { // 是否不足1天, 有几天, 减去天数后的时间差
	if num = sinter / tt; num > 0 {
		return num, sinter - (num * tt)
	} else {
		return 0, sinter
	}
	return
}
