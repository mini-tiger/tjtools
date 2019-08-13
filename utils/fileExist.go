package utils

import "os"

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
