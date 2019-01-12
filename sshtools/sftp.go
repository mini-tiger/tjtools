package sshtools

import (
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"path"
	"time"
)

//var (
//	sftpClient *sftp.Client
//)

func SftpUpload(host, passwd, sFile, dFile string) (err error) {

	//start := time.Now()
	sftpClient, err := connect("root", passwd, host, 22) //远程链接打开
	if err != nil {
		//log.Fatal(err)
		return
	}
	defer sftpClient.Close()

	_, err = sftpClient.Stat(path.Dir(dFile)) //远程目录确认状态是否存在
	if err != nil {
		//log.Fatal("/tmp/" + " remote path not exists!")\
		err = errors.New(fmt.Sprintf(" remote path %s not exists,err:%s", path.Dir(dFile), err))
		return
	}

	//backupDirs, err := ioutil.ReadDir("c:\\1.log")
	//if err != nil {
	//	log.Fatal("c:\\1.log  local path not exists!")
	//}
	//uploadFile(sftpClient, "D:\\work\\project-dev\\src\\godev\\mymodels\\ssh并发运行脚本\\1.sh", "/tmp/1.sh")
	return uploadFile(sftpClient, sFile, dFile)

}
func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

func uploadFile(sftpClient *sftp.Client, localFilePath string, remotePath string) (err error) {
	srcFile, err := os.Open(localFilePath) //打开需要上传的本地文件
	if err != nil {
		err = errors.New(fmt.Sprintf("os.Open file:%s error : %s", localFilePath, err))
		return
	}
	defer srcFile.Close()

	//var remoteFileName = path.Base(localFilePath)

	dstFile, err := sftpClient.Create(remotePath) //创建远程文件
	if err != nil {
		//fmt.Println("sftpClient.Create error : ", remotePath)
		err = errors.New(fmt.Sprintf("sftpClient.Create file:%s error : %s", remotePath, err))
		return

	}
	defer dstFile.Close()

	ff, err := ioutil.ReadAll(srcFile) // 本地文件内容数据
	if err != nil {
		//fmt.Println("ReadAll error : ", localFilePath)
		err = errors.New(fmt.Sprintf("ioutils.ReadAll:%s error : %s", localFilePath, err))
		return
	}
	dstFile.Write(ff) //远程文件 写入 本地文件内容
	//fmt.Println(localFilePath + "  copy file to remote server finished!")
	return nil
}
