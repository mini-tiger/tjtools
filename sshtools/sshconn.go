package sshtools

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type SShInfo struct {
	IP, Username string
	Passwd       string
	Port         int
	client       *ssh.Client
	Session      *ssh.Session
	Result       string
}

func New_ssh(port int, args ...string) *SShInfo {
	temp := new(SShInfo)
	temp.Port = port
	temp.IP = args[0]
	temp.Username = args[1]
	temp.Passwd = args[2]
	return temp

}
func (cli *SShInfo) Connect() error {
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(cli.Passwd))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig := &ssh.ClientConfig{
		User:            cli.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr := fmt.Sprintf("%s:%d", cli.IP, cli.Port)

	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return err
	}

	// create session
	session, err := client.NewSession()
	if err != nil {
		defer cli.Close_session()
		return err
	}
	cli.Session = session
	return nil
}
func (cli *SShInfo) Close_session() {
	cli.Session.Close()
}

func (cli *SShInfo) TerminalRun(cmd, dFile string, shellFile bool) (err error, outResult string) {
	w, err := cli.Session.StdinPipe()
	if err != nil {
		//panic(err)
		return
	}
	r, err := cli.Session.StdoutPipe()
	if err != nil {
		//panic(err)
		return
	}
	e, err := cli.Session.StderrPipe()
	if err != nil {
		//panic(err)
		return
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal 建立终端
	if err = cli.Session.RequestPty("vt100", 40, 80, modes); err != nil { //term:xerm 是彩色显示
		//log.Fatal("request for pseudo terminal failed: ", err)
		return
	}

	in, out := MuxShell(w, r, e)
	if err = cli.Session.Shell(); err != nil { //打开仿真shell
		//log.Fatal(err)
		return
	}
	//<-out 通信out第一次返回的是 linux 登录信息,可以跳过

	if shellFile { // 如果是脚本
		//in <- fmt.Sprintf("chmod 777 %s", dFile)
		in <- "chmod 777 /tmp/1.sh"
	}
	in <- "bash /tmp/1.sh"
	//fmt.Println(fmt.Sprintf("chmod 777 %s", dFile))

	in <- "exit" //todo 需要用这条 关闭Session

	//fmt.Println(cmd)

	for {
		if k, ok := <-out; ok {
			outResult += fmt.Sprintf("%s", k) //todo 所有out通道中记录的 返回信息 打开出来
			//fmt.Println(k)
		} else {
			break
		}
	}

	cli.Session.Close()
	cli.Close_session()
	//session.Close()
	//session.Wait()
	return
}

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 0)
	out := make(chan string, 4)
	var wg sync.WaitGroup
	wg.Add(1) //shell 退出前，Shell的进程
	go func() {
		for cmd := range in { //todo in通道中 所有 需要执行的命令 依次执行
			wg.Add(1)
			w.Write([]byte(cmd + "\n")) //w 是管道输入
			wg.Wait()                   //等待命令完成
		}
	}()

	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:]) //todo 标准输出管道的 内容，stdoutpipe,是io.Reader接口有reader方法，将传入的[]byte 写入
			if err != nil {
				if err == io.EOF { //如果EOF 退出
					//fmt.Println("exit")
					//wg.Done()
				}
				//fmt.Println(err.Error())
				close(in)
				close(out)
				return
			}
			t += n //每次命令结果 追加至buf
			result := string(buf[:t])
			if strings.Contains(result, "password:") ||
				strings.Contains(result, "#") { //匹配是否执行完成
				out <- result
				t = 0 //t是临时存 当前命令返回的结果，清空
				wg.Done()
			}
		}
	}()
	return in, out
}
