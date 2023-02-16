package message

import (
	"fmt"
	"io"
	"net"
	"sync"
	"tiktok/config"
)

var chatConnMap = sync.Map{}

func InitMessageServer() {

	return
	listen, err := net.Listen("tcp", config.MessageServerIp)
	if err != nil {
		fmt.Printf("Run MEssage server Failed:%v\n", err)
		return
	}
	fmt.Printf("is listening addr %s ..........\n", config.MessageServerIp)

	for {
		conn, err := listen.Accept()
		fmt.Println("new conn is accept.................... client network:", conn.RemoteAddr().String())
		if err != nil {
			fmt.Printf("package message [func-InitMessageServer] Accept error:%v\n", err)
			continue
		}
		// go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("read message err:%v\n", err)
			continue
		}
	}

	// var event = common.MessageActionResponse{}
}
