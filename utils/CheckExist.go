package utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Exist(file string) bool {

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			//fmt.Printf("文件: %s 不存在\n", file)
			return false
		}
	} else {
		//fmt.Printf("文件: %s 存在\n", file)
		return true
	}

	return false
}
