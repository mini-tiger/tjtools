package file

import (
	"errors"
	"gitee.com/taojun319/tjtools/file"
	"os"
	"strings"

)

func GetFatherDir(sf string) (string, error) {
	if !file.IsExist(sf) {
		return "", errors.New(sf + "Not Exist")
	}
	dirList := strings.Split(sf, string(os.PathSeparator))
	if len(dirList) <= 1 {
		return sf, nil
	}
	return strings.Join(dirList[0:len(dirList)-1], string(os.PathSeparator)), nil
}
