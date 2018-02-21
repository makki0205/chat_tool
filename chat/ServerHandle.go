package chat

import (
	"net"
	"fmt"
)

func ServerHandle(conn net.Conn){
	socket := SocketManager(conn)
	for{
		message := socket.Read()
		if message == ""{
			break
		}
		fmt.Println("[受信]: " + message)
		SendAll(message)
	}
}
func SendAll(message string){
	for _, socket := range sockets {
		socket.Send(message)
	}
}