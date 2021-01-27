package file

import (
	"fmt"
	"os"
	"syscall"
)

func CheckDirPerm(filePath string) (err error) {

	var file os.FileInfo
	file, err = os.Stat(filePath)
	if err != nil {
		return
	}

	s := file.Sys().(*syscall.Stat_t)

	if err = os.Chown(filePath, int(s.Uid), int(s.Gid)); err != nil {
		return
	}
	return nil
}
func main() {
	fmt.Println(CheckDirPerm("/root"))
}
