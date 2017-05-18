package Chat

import (
	"fmt"
	"net"
)

func ServerInit(){
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Accept error: %s\n", err)
		return
	}
	sock := SocketManager(conn)
	go ReadHandle(sock)

	SendHandle(sock)
}

