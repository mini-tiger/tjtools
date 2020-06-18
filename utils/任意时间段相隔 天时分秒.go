package utils

//import (
//	"fmt"
//	"time"
//)

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
const daySec=86400
var Sl *sync.RWMutex = new(sync.RWMutex)
type TimeCount struct {
	Day, Interval, H, M, S uint64
}

var TimeCountFree = sync.Pool{
	New: func() interface{} {
		return &TimeCount{}
	},
}


func (t *TimeCount) GetCurrentTime() *TimeCount {
	Sl.Lock()
	defer Sl.Unlock()
	return t
}

func (t *TimeCount) ComputeTime(sinter int64) {
	t.getBasic(uint64(sinter))

	t.oneDay()

	//s = uint64(sinter) - (Day * 86400) - (h * 3600) - (m * 60)
	atomic.StoreUint64(&t.S,uint64(sinter) - (t.Day * 86400) - (t.H * 3600) - (t.M * 60))

}

func (t *TimeCount) oneDay() { // 一天内的小时分钟
	atomic.StoreUint64(&t.H,t.Interval / 3600)
	//t.H = t.Interval / 3600
	if t.H == 0 {
		//t.M = t.Interval / 60
		atomic.StoreUint64(&t.M,t.Interval / 60)
	} else {
		//t.Interval = t.Interval - (3600 * t.H)
		atomic.StoreUint64(&t.Interval,t.Interval - (3600 * t.H))
		//t.M = t.Interval / 60
		atomic.StoreUint64(&t.M,t.Interval / 60)
	}

}

func  (t *TimeCount)getBasic(sinter uint64) { // 是否不足1天, 有几天, 减去天数后的时间差
	if t.Day = sinter / daySec; t.Day > 0 {
		//interval = sinter - (Day * tt)
		atomic.StoreUint64(&t.Interval,sinter - (t.Day * daySec))
	} else {
		//Day = 0
		atomic.StoreUint64(&t.Day,0)
		//interval = sinter
		atomic.StoreUint64(&t.Interval,sinter)
	}

}


