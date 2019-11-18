package file

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func ReadLine(r *bufio.Reader) ([]byte, error) {
	line, isPrefix, err := r.ReadLine()
	for isPrefix && err == nil {
		var bs []byte
		bs, isPrefix, err = r.ReadLine()
		line = append(line, bs...)
	}

	return line, err
}
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

func GetFatherDir(sf string) (string, error) {
	if !IsExist(sf) {
		return "", errors.New(sf + "Not Exist")
	}
	dirList := strings.Split(sf, string(os.PathSeparator))
	if len(dirList) <= 1 {
		return sf, nil
	}
	return strings.Join(dirList[0:len(dirList)-1], string(os.PathSeparator)), nil
}
