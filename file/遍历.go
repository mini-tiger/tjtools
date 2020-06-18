package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

//func GetFileList(path string) ([]string, error) {
//	files := make([]string, 0)
//
//	f, err := os.Stat(path)
//	if err != nil {
//		return files, errors.New(fmt.Sprintf("path: %s ,Err:%s", path, err))
//	}
//	if !f.IsDir() {
//		return files, errors.New(fmt.Sprintf("path: %s ,Not Dir", path))
//	}
//
//	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
//		if f == nil {
//			return err
//		}
//		if f.IsDir() {
//			return nil
//		}
//		files = append(files, path)
//		return nil
//	})
//	if err != nil {
//		return files, errors.New(fmt.Sprintf("path: %s ,err:%s", path, err))
//	}
//	return files, nil
//}



var FSL = new(sync.RWMutex)
type SelectFiles struct {
	Files []string
	Path string
}
var SelectFilesFree = sync.Pool{
	New: func() interface{} {
		return &SelectFiles{}
	},
}
//var files []string = make([]string, 0, 1<<13) // 不超过8192,可以重复利用
//var syncFile *sync.Mutex = new(sync.Mutex)

func (s *SelectFiles)GetFiles() []string {
	FSL.RLock()
	defer FSL.RUnlock()

	return s.Files
}
func (s *SelectFiles)Clearfiles() {
	FSL.Lock()
	defer FSL.Unlock()
	s.Files = s.Files[0:0]
}

func (s *SelectFiles)Cleanfiles() {
	FSL.Lock()
	defer FSL.Unlock()
	s.Files = nil
}

func  (s *SelectFiles)GetFileList() (err error) {
	//files := make([]string, 0)
	//tfiles := filesPool.Get().([]string)

	f, err := os.Stat(s.Path)
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,Err:%s", s.Path, err))
		return err
	}
	if !f.IsDir() {
		err = errors.New(fmt.Sprintf("path: %s ,Not Dir", s.Path))
		return err
	}
	FSL.Lock()
	defer FSL.Unlock()

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

	//fmt.Printf("%+v,%d\n",tfiles,len(tfiles))
	//if len(files) > 0 {
	//	return err
	//}
	//fmt.Printf("%p\n",tfiles)
	return nil
}