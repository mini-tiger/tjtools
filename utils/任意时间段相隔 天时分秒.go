package utils

import (
	"fmt"
	"time"
)

func main() {
	timeLayout := "2006-01-02 15:04:05"
	start := time.Now().Unix()
	start = 1576899615 -86401
	fmt.Println(time.Unix(start, 0).Format(timeLayout))
	//time.Sleep(time.Duration(5)*time.Second)
	endtime := time.Now().Unix()
	sinter := endtime - start
	var d, h, m, s int64
	GetTime(&sinter, &d, &h, &m, &s)
	fmt.Printf("%d天%d小时%d分钟%d秒", d, h, m, s)
}

func GetTime(sinter, d, h, m, s *int64) {
	interval := getBasic(sinter, d,86400)
	oneDay(&interval,h,m)



	*s = *sinter - (*d * 86400) - (*h * 3600) - (*m * 60)
	return
}

func oneDay(interval, h,m *int64) { // 一天内的小时分钟
	*h = *interval / 3600
	if *h== 0 {
		*m = *interval / 60
	} else {
		*interval = *interval - (3600 * (*h))
		*m = *interval / 60
	}
}

func getBasic(sinter,day *int64, tt int64) (interval int64) { // 是否不足1天, 有几天, 减去天数后的时间差
	if *day = *sinter / tt; *day > 0 {
		return *sinter - (*day * tt)
	} else {
		*day=0
		return *sinter
	}

}
