package udp

import (
	"net"
)

func main() {

	SendUdp("127.0.0.1:8000", []byte("this is test"))
}
func SendUdp(udpAddr string, sendData []byte) (err error) {
	addr, err := net.ResolveUDPAddr("udp", udpAddr)
	if err != nil {
		//fmt.Println("Can't resolve address: ", err)
		//os.Exit(1)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		//fmt.Println("Can't dial: ", err)
		//os.Exit(1)
		return
	}
	defer conn.Close()
	_, err = conn.Write(sendData) // todo 发送到server
	if err != nil {
		//fmt.Println("failed:", err)
		//os.Exit(1)
		return
	}
	return nil
	//data := make([]byte, 4)
	//_, err = conn.Read(data) //todo 读取server发回来的
	//if err != nil {
	//	fmt.Println("failed to read UDP msg because of ", err)
	//	os.Exit(1)
	//}
	//t := binary.BigEndian.Uint32(data)
	//fmt.Println(time.Unix(int64(t), 0).String())
	//os.Exit(0)
}
