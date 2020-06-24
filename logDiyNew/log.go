package logDiyNew

import (
	"fmt"
	nxlog "github.com/ccpaging/nxlog4go"
	"io"
	"os"
	"time"
)

var (
	Logge *nxlog.Logger
)

func WLog(str string) { // 在配置文件没有加载，日志方法没有生效前，写入日志
	f, err1 := os.OpenFile("run.log", os.O_CREATE|os.O_SYNC|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err1 != nil {
		_, _ = os.Stdout.Write([]byte(fmt.Sprintf("wLog file err:%s\n", err1)))
	}
	t1 := time.Now()
	//fmt.Println(t1.Format("2006 01-02 15:04:05"))

	str = fmt.Sprintf(" [%s] [%s] (%s) \n", t1.Format("2006 01-02 15:04:05"), "ERROR", str)
	_, _ = f.Write([]byte(str))

}

func InitLog1(logfile string, maxDays int, color bool, level string, std bool) *nxlog.Logger {
	nxlog.FileFlushDefault = 5

	rfw := nxlog.NewRotateFileWriter(logfile, true)
	_ = rfw.SetOption("daily", true)
	_ = rfw.SetOption("maxbackup", maxDays)

	var ww io.Writer
	if std {
		ww = io.MultiWriter(os.Stdout, rfw) //todo 同时输出到rfw 与 系统输出
	} else {
		ww = rfw //todo 同时输出到rfw 与 系统输出
	}

	// Get a new logger instance
	// todo FINEST 级别最低
	// todo %P prefix, %N 行号
	Logge = nxlog.New(os.Stdout, "", 7)
	Logge.SetOutput(ww)
	_ = Logge.SetOption("caller", true) //runtime.caller(3)  日志触发上报的层级
	_ = Logge.SetOption("color", color)
	_ = Logge.SetOption("level", level)
	//Logge.SetOption("prefix","this is prefix")
	//Logge.SetLayout(nxlog.NewPatternLayout("%P %Y %T [%L] (%s LineNo:%N) %M"))
	Logge.SetLayout(nxlog.NewPatternLayout("%Y %T [%L] (%s LineNo:%N) %M"))

	return Logge

}

//func Logger() *Log1 {
//	lock.RLock()
//	defer lock.RUnlock()
//	return logger
//}

//func main() {
//	//var Log *Log1
//
//	InitLog1("C:\\work\\go-dev\\AutoNomy\\nxlogNew\\1.log", 3)
//	//Log = Logger()
//	logge.Printf("%d\n", 1111)
//	logge.Printf("%d\n", 1111)
//	logge.Error("%d\n", 1111)
//	logge.Debug("%d\n", 1111)
//}
