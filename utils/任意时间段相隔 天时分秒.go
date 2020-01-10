package utils

import (
	"fmt"
	"time"
)

//func main() {
//	timeLayout := "2006-01-02 15:04:05"
//	start := time.Now().Unix()
//	start = 1576899615 - 86401
//	fmt.Println(time.Unix(start, 0).Format(timeLayout))
//	//time.Sleep(time.Duration(5)*time.Second)
//	endtime := time.Now().Unix()
//	sinter := endtime - start
//	var d, h, m, s int64
//	d, h, m, s = GetTime(sinter)
//	fmt.Printf("%d天%d小时%d分钟%d秒", d, h, m, s)
//}
var Day, interval, num ,h,m,s uint64


func GetCurrentTime() (*uint64, *uint64, *uint64, *uint64) {
	return &Day, &h, &m, &s
}

func GetTime(sinter int64) {
	getBasic(uint64(sinter), 86400)

	oneDay()

	s = uint64(sinter) - (Day * 86400) - (h * 3600) - (m * 60)

}

func oneDay()  { // 一天内的小时分钟
	h = interval / 3600
	if h == 0 {
		m = interval / 60
	} else {
		interval = interval - (3600 * h)
		m = interval / 60
	}

}

func getBasic(sinter uint64, tt uint64) { // 是否不足1天, 有几天, 减去天数后的时间差
	if num = sinter / tt; num > 0 {
		interval = sinter - (num * tt)
	} else {
		num = 0
		interval = sinter
	}

}
