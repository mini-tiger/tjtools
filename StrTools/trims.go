package StrTools

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"bytes"
	"strings"
	"regexp"
)


// StrTrims 清除字符串中 空行
// DeleteBlankFile 清除文件中 空行


func DeleteBlankFile(srcFilePah string, destFilePath string) error {
	srcFile, err := os.OpenFile(srcFilePah, os.O_RDONLY, 0666)
	defer srcFile.Close()
	if err != nil {
		return err
	}
	srcReader := bufio.NewReader(srcFile)
	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer destFile.Close()
	if err != nil {
		return err
	}

	for {
		str, _ := srcReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print("The file end is touched.")
				break
			} else {
				return err
			}
		}
		if 0 == len(str) || str == "\r\n" {
			continue
		}
		fmt.Print(str)
		destFile.WriteString(str)
	}
	return nil
}

func main() {
	//DeleteBlankFile("c:\\1.txt", "c:\\dest.txt")
	//destbyte:=

}

func StrTrims(ss string) (d string,e error)  {
	dest := bytes.NewBuffer([]byte{})
	ssb := bytes.NewBuffer([]byte(ss))
	for {
		str, err := ssb.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//fmt.Print("The file end is touched.")
				break
			} else {
				//fmt.Println(err)
				e=err
				return
			}
		}
		str = strings.Replace(str, string(byte(13)), "", -1) // 去除所有win 回车
		str = strings.Replace(str, string(byte(8)), "", -1)  //去除所有退格

		//if 0 == len(str) || str == "\r\n" {
		//	continue
		//}
		if 1 == len(str) && (str == "\n" || str == "\r") {
			continue
		}

		re := regexp.MustCompile("\\s{2,}?")
		str = re.ReplaceAllString(str, "")

		sl := strings.Count(str, string(byte(10))) //超过2 个换行，保留一个，其它删除

		if sl < 1 {
			str = str + string(byte(10))
		}
		//fmt.Println(str)
		//fmt.Println(1)
		//fmt.Println([]byte(str))
		dest.WriteString(str)
	}
	return dest.String(),nil
}