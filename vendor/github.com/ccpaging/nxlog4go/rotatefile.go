// Copyright (C) 2017, ccpaging <ccpaging@gmail.com>.  All rights reserved.

package nxlog4go

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

// Rotate File buffer writer
type RotateFileWriter struct {
	sync.RWMutex
	// The opened file buffer writer
	*FileBufWriter
	// File header/trailer
	header, footer string
	// Rotate at size
	maxsize int
	// Rotate daily
	daily   bool
	curtime time.Time
	// Keep old files (.1, .2, etc)
	maxbackup int
}

// NewRotateFileWriter creates rotator which writes to the file buffer writer
func NewRotateFileWriter(path string) *RotateFileWriter {
	rfw := &RotateFileWriter{
		FileBufWriter: NewFileBufWriter(path),
		daily:         false,
	}
	fi, err := rfw.FileBufWriter.Stat()
	if err != nil {
		rfw.curtime = time.Now()
	} else {
		rfw.curtime = fi.ModTime()
	}
	return rfw
}

// Write binaries to the file.
// It will rotate files if necessary
func (rfw *RotateFileWriter) Write(bb []byte) (n int, err error) {
	rfw.Lock()
	defer rfw.Unlock()

	if (rfw.maxsize > 0 && rfw.Size() >= rfw.maxsize) ||
		(rfw.daily && rfw.curtime.Day() != time.Now().Day()) {

		rfw.Rotate()
	}

	if len(rfw.header) > 0 && rfw.Size() == 0 {
		layout := NewPatternLayout(rfw.header)
		rfw.FileBufWriter.Write(layout.Format(&LogRecord{Created: time.Now()}))
	}

	return rfw.FileBufWriter.Write(bb)
}

// Rename history log files to "<name>.???.<ext>"
func intBackup(newName string, pattern string, backup int) {
	// May compress new log file here

	LogLogTrace("Backup %s", newName)

	ext := path.Ext(pattern)                   // like ".log"
	path := pattern[0 : len(pattern)-len(ext)] // include dir

	// May create backup directory here

	var (
		n    int
		err  error
		slot string
	)
	for n = 1; n <= backup; n++ {
		slot = path + fmt.Sprintf(".%d", n) + ext
		_, err = os.Stat(slot)
		if err != nil {
			break
		}
	}
	if err == nil {
		// Full. Remove last
		os.Remove(slot)
		n--
	}

	// May compress previous log file here

	for ; n > 1; n-- {
		prev := path + fmt.Sprintf(".%d", n-1) + ext
		LogLogTrace("Rename %s to %s", prev, slot)
		os.Rename(prev, slot)
		slot = prev
	}

	LogLogTrace("Rename %s to %s", newName, path+".1"+ext)
	os.Rename(newName, path+".1"+ext)
}

func (rfw *RotateFileWriter) Rotate() {
	defer func() {
		rfw.curtime = time.Now()
	}()

	// Write footer
	if len(rfw.footer) > 0 && rfw.Size() > 0 {
		layout := NewPatternLayout(rfw.footer)
		rfw.FileBufWriter.Write(layout.Format(&LogRecord{Created: time.Now()}))
	}

	LogLogTrace("Close %s", rfw.FileBufWriter.Name())
	rfw.FileBufWriter.Close()

	name := rfw.FileBufWriter.Name()
	if rfw.maxbackup <= 0 {
		os.Remove(name)
		return
	}

	// File existed. File size > maxsize. Rotate
	newName := name + time.Now().Format(".20060102-150405")
	LogLogTrace("Rename %s to %s", name, newName)
	err := os.Rename(name, newName)
	if err != nil {
		LogLogError(err)
		return
	}

	intBackup(newName, name, rfw.maxbackup)
}

// Set the file header(chainable).  Must be called before the first log
// message is written.
func (rfw *RotateFileWriter) SetFileName(path string) *RotateFileWriter {
	rfw.Lock()
	defer rfw.Unlock()
	LogLogTrace("Set file name as %s", path)
	rfw.FileBufWriter.Close()
	rfw.FileBufWriter = NewFileBufWriter(path)
	return rfw
}

// Set the file header(chainable).  Must be called before the first log
// message is written.  These are formatted similar to the FormatLogRecord (e.g.
// you can use %D and %T in your header for date and time).
func (rfw *RotateFileWriter) SetHead(header string) *RotateFileWriter {
	rfw.Lock()
	defer rfw.Unlock()
	rfw.header = header
	return rfw
}

// Set the file footer (chainable).  Must be called before the first log
// message is written.  These are formatted similar to the FormatLogRecord (e.g.
// you can use %D and %T in your footer for date and time).
func (rfw *RotateFileWriter) SetFoot(footer string) *RotateFileWriter {
	rfw.Lock()
	defer rfw.Unlock()

	rfw.footer = footer
	return rfw
}

// Set rotate at size (chainable). Must be called before the first log message
// is written.
func (rfw *RotateFileWriter) SetMaxSize(maxsize int) *RotateFileWriter {
	rfw.Lock()
	defer rfw.Unlock()

	rfw.maxsize = maxsize
	return rfw
}

// Set rotate daily (chainable). Must be called before the first log message is
// written.
func (rfw *RotateFileWriter) SetDaily(daily bool) *RotateFileWriter {
	rfw.Lock()
	defer rfw.Unlock()

	rfw.daily = daily
	return rfw
}

// Set max backup files. Must be called before the first log message
// is written.
func (rfw *RotateFileWriter) SetMaxBackup(maxbackup int) *RotateFileWriter {
	rfw.Lock()
	defer rfw.Unlock()

	rfw.maxbackup = maxbackup
	return rfw
}

func (rfw RotateFileWriter) GetMaxBackup() int {
	return rfw.maxbackup
}
