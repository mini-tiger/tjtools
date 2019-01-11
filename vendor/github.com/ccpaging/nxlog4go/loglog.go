// Copyright (C) 2017, ccpaging <ccpaging@gmail.com>.  All rights reserved.

package nxlog4go

var loglog *Logger = nil

// Return internal logger.
// This logger used to output log statements from within the package.
// Do not set any filters.
func GetLogLog() *Logger {
	if loglog == nil {
		loglog = New(DEBUG).SetPrefix("lg4g").SetPattern("%T %P %L %M\n").SetCaller(false)
	}
	return loglog
}

func LogLogDebug(arg0 interface{}, args ...interface{}) {
	if loglog == nil {
		return
	}
	if loglog.skip(DEBUG) {
		return
	}
	loglog.intLog(DEBUG, intMsg(arg0, args...))
}

func LogLogTrace(arg0 interface{}, args ...interface{}) {
	if loglog == nil {
		return
	}
	if loglog.skip(TRACE) {
		return
	}
	loglog.intLog(TRACE, intMsg(arg0, args...))
}

func LogLogInfo(arg0 interface{}, args ...interface{}) {
	if loglog == nil {
		return
	}
	if loglog.skip(INFO) {
		return
	}
	loglog.intLog(INFO, intMsg(arg0, args...))
}

func LogLogWarn(arg0 interface{}, args ...interface{}) {
	if loglog == nil {
		return
	}
	if loglog.skip(WARNING) {
		return
	}
	loglog.intLog(WARNING, intMsg(arg0, args...))
}

func LogLogError(arg0 interface{}, args ...interface{}) {
	if loglog == nil {
		return
	}
	if loglog.skip(ERROR) {
		return
	}
	loglog.intLog(ERROR, intMsg(arg0, args...))
}
