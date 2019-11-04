package logDiyNew

import (
	"fmt"
	nxlog "github.com/ccpaging/nxlog4go"
	"io"
	"os"
	"time"
)

var (
	logge *nxlog.Logger
)
//var logger *Log1
//var lock = new(sync.RWMutex)

//type Log1 struct{}

//func (l *Log1) Printf(arg0 ...interface{}) {
//	logge.Info(arg0...)
//}
//func (l *Log1) Error(arg0 ...interface{}) error {
//	logge.SetOption("head", "[error]")
//	return logge.Error(arg0...)
//}
//
////
//func (l *Log1) Debug(arg0 ...interface{}) {
//	logge.Debug(arg0...)
//}
//
//func (l *Log1) Fatalln(m ...interface{}) {
//	logge.Error(m)
//	os.Exit(1)
//}
//func (l *Log1) Fatalf(arg0 interface{}, args ...interface{}) {
//	logge.Error(arg0, args)
//	os.Exit(1)
//}

func WLog(str string) { // 在配置文件没有加载，日志方法没有生效前，写入日志
	f, err1 := os.OpenFile("run.log", os.O_CREATE|os.O_SYNC|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err1 != nil {
		os.Stdout.Write([]byte(fmt.Sprintf("wLog file err:%s\n", err1)))
	}
	t1 := time.Now()
	//fmt.Println(t1.Format("2006 01-02 15:04:05"))

	str = fmt.Sprintf(" [%s] [%s] (%s) \n", t1.Format("2006 01-02 15:04:05"), "ERROR", str)
	f.Write([]byte(str))

}

func InitLog1(logfile string, maxDays int) *nxlog.Logger {
	//fileName := Config().Logfile
	fileName := logfile
	//logFile, err := os.Create(fileName)
	//if err != nil {
	//	log.Fatalln("open file error !")
	//}

	nxlog.FileFlushDefault = 5 // 修改默认写入硬盘时间
	//nxlog.LogCallerDepth = 3                                                        //runtime.caller(3)  日志触发上报的层级
	rfw := nxlog.NewRotateFileWriter(fileName, true)
	rfw.SetOption("daily", true);
	rfw.SetOption("maxbackup", maxDays);
	var ww io.Writer

	ww = io.MultiWriter(os.Stdout, rfw) //todo 同时输出到rfw 与 系统输出

	//s:=fmt.Sprintf("test 111,%d\n",22222)
	//sb:=bytes.NewBufferString(s)
	//ss.Write(sb.Bytes())

	// Get a new logger instance
	// todo FINEST 级别最低
	// todo %p prefix, %N 行号
	logge = nxlog.New(os.Stdout, "", 7)
	logge.SetOutput(ww)
	logge.SetOption("prefix","this is prefix")
	logge.SetLayout(nxlog.NewPatternLayout("%P; %Y %T [%L] (%s LineNo:%N) %M"))
	//logge.Info("read config file ,successfully") // 走到这里代表配置文件已经读取成功
	//logge.Info("日志文件最多保存%d天", Config().LogMaxDays)
	//logge.Info("logging on %s", fileName)
	//logge.Info("进程已启动, 当前进程PID:%d\n", os.Getpid())
	return logge

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
