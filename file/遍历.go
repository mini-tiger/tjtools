package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetFileList(path string) ([]string, error) {
	files := make([]string, 0)

	f, err := os.Stat(path)
	if err != nil {
		return files, errors.New(fmt.Sprintf("path: %s ,Err:%s", path, err))
	}
	if !f.IsDir() {
		return files, errors.New(fmt.Sprintf("path: %s ,Not Dir", path))
	}

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
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
		return files, errors.New(fmt.Sprintf("path: %s ,err:%s", path, err))
	}
	return files, nil
}
