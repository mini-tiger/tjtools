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



var files []string = make([]string, 0, 1<<13) // 不超过8192,可以重复利用
var syncFile *sync.Mutex = new(sync.Mutex)

func GetFiles() []string {
	syncFile.Lock()
	defer func() {
		syncFile.Unlock()
	}()
	return files
}
func Clearfiles() {
	syncFile.Lock()
	defer func() {
		syncFile.Unlock()
	}()
	files = files[0:0]
}

func GetFileList(path *string) (err error) {
	//files := make([]string, 0)
	//tfiles := filesPool.Get().([]string)

	f, err := os.Stat(*path)
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,Err:%s", *path, err))
		return err
	}
	if !f.IsDir() {
		err = errors.New(fmt.Sprintf("path: %s ,Not Dir", *path))
		return err
	}
	syncFile.Lock()
	defer func() {
		syncFile.Unlock()
	}()

	err = filepath.Walk(*path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,err:%s", *path, err))
		return err
	}

	//fmt.Printf("%+v,%d\n",tfiles,len(tfiles))
	//if len(files) > 0 {
	//	return err
	//}
	//fmt.Printf("%p\n",tfiles)
	return nil
}