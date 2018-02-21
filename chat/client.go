package chat

import (
	"fmt"
	"net"
)

func ClientInit(ip string){
	conn, err := net.Dial("tcp", ip + ":1234")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	defer conn.Close()
	sock := SocketManager(conn)
	go ReadHandle(sock)
	SendHandle(sock)
}