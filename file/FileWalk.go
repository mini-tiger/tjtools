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
//
//func main() {
//	for{
//
//		aa:=FilesFree.Get().(*Files)
//		//aa:=NewFilesStruct("/home/go/GoDevEach/works/haifei/syncHtml_test8/htmlData/")
//		aa.Path="/home/go/GoDevEach/works/haifei/syncHtml_test8/htmlData/"
//		_ = aa.ScanningFiles()
//		FilesFree.Put(aa)
//		fmt.Printf("%d,%p,%p\n",aa.Size(),aa,aa.FileAbs)
//		aa=nil
//		time.Sleep(1*time.Second)
//	}
//
//	//for {
//	//
//	//	path := "/home/go/GoDevEach/works/haifei/syncHtml_test8/htmlData/"
//	//	_ = file.GetFileList(&path) // 遍历目录 包括子目录
//	//	fmt.Printf("%d,%p\n", len(file.GetFiles()), file.GetFiles())
//	//
//	//
//	//	if len(file.GetFiles()) > 0 {
//	//		//var h = OnceHtmlFusionFiles.Get().(*modules.HtmlFusionFiles)
//	//		//debug.SetGCPercent(100)
//	//
//	//		file.Clearfiles()
//	//
//	//	}
//	//	fmt.Printf("%d,%p\n", len(file.GetFiles()), file.GetFiles())
//	//
//	//	//}
//	//	time.Sleep(time.Duration(1 * time.Second))
//	//}
//
//}

type Files struct {
	sync.RWMutex
	FileAbs []string
	Path    string
}

var FilesFree = sync.Pool{
	New: func() interface{} {
		return &Files{}
	},
}

//var files []string = make([]string, 0, 1<<13) // xxx 不超过8192,可以重复利用
//var syncFile *sync.Mutex = new(sync.Mutex)

func NewFilesStruct(path string) *Files {
	return &Files{
		FileAbs: make([]string, 0),
		Path:    path,
	}
}

func (fi *Files) GetFiles() []string {
	fi.Lock()
	defer func() {
		fi.Unlock()
	}()
	return fi.FileAbs
}

func (fi *Files) Size() int {
	fi.RLock()
	//len := len(this.M)
	defer func() {
		fi.RUnlock()
	}()
	return len(fi.FileAbs)
}

func (fi *Files) IsEmpty() bool {

	return fi.Size() == 0
}

func (fi *Files) ClearFiles() {
	fi.Lock()
	defer func() {
		fi.Unlock()
	}()
	fi.FileAbs = fi.FileAbs[0:0]
}

func (fi *Files) ScanningFiles() (err error) {

	fi.ClearFiles()
	ff, err := os.Stat(fi.Path)
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,Err:%s", fi.Path, err))
		return err
	}
	if !ff.IsDir() {
		err = errors.New(fmt.Sprintf("path: %s ,Not Dir", fi.Path))
		return err
	}
	fi.Lock()
	defer func() {
		fi.Unlock()
	}()

	err = filepath.Walk(fi.Path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fi.FileAbs = append(fi.FileAbs, path)
		return nil
	})
	if err != nil {
		err = errors.New(fmt.Sprintf("path: %s ,err:%s", fi.Path, err))
		return err
	}

	//fmt.Printf("%+v,%d\n",tfiles,len(tfiles))
	//if len(files) > 0 {
	//	return err
	//}
	//fmt.Printf("%p\n",tfiles)
	return nil
}
