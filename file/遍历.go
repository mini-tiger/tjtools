package file

import (
	"errors"
	"fmt"
	//nsema "gitee.com/taojun319/tjtools/control"
	"os"
	"path/filepath"
	"sync"
	//"time"
)

//var FSL = new(sync.RWMutex)

type SelectFiles struct {
	sync.RWMutex
	Files []string
	Path  string
}

var SelectFilesFree = sync.Pool{
	New: func() interface{} {
		return &SelectFiles{}
	},
}

//var files []string = make([]string, 0, 1<<13) // 不超过8192,可以重复利用
//var syncFile *sync.Mutex = new(sync.Mutex)

func (s *SelectFiles) GetFiles() []string {
	s.RLock()
	defer s.RUnlock()

	return s.Files
}
func (s *SelectFiles) Clearfiles() {
	s.Lock()
	defer s.Unlock()
	s.Files = s.Files[0:0]
}

func (s *SelectFiles) Cleanfiles() {
	s.Lock()
	defer s.Unlock()
	s.Files = nil
}
func (s *SelectFiles) Len() uint64 {
	return uint64(len(s.Files))
}

func (s *SelectFiles) GetFileList() (err error) {
	f, err := os.Stat(s.Path)
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,Err:%s", s.Path, err))
		return err
	}
	if !f.IsDir() {
		err = errors.New(fmt.Sprintf("path: %s ,Not Dir", s.Path))
		return err
	}
	s.Lock()
	defer s.Unlock()
	s.Files = make([]string, 0)
	err = filepath.Walk(s.Path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		s.Files = append(s.Files, path)
		return nil
	})
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,err:%s", s.Path, err))
		return err
	}

	return nil
}

//
//var FilesChan1 chan *SelectFiles = make(chan *SelectFiles, 0)
//var sema2 *nsema.Semaphore = nsema.NewSemaphore(2)
//
//
//func main() {
//
//
//	go Revice2()
//	go Push2()
//
//	select {}
//}
//
//func push22()  {
//	rr := new(SelectFiles)
//	rr.Path="/home/go/src/godev/内存"
//	_ = rr.GetFileList()
//	FilesChan1<-rr
//}
//
//func Push2() {
//
//	for {
//		push22()
//		time.Sleep(10*time.Second)
//	}
//
//}
//
//func Revice2() {
//	for {
//		select {
//		case sa2 := <-FilesChan1:
//			sema2.Acquire()
//			sa2.GetFiles()
//			//Clear(sa2)
//			sema2.Release()
//		}
//	}
//}
