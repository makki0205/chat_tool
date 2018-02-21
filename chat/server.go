package chat


import (
	"fmt"
	"net"
	"bufio"
	"os"
)
var sockets []*Socket

func ServerInit() {
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	go AcceptHandle(listener)

	ServerSendHandle()
}
func AcceptHandle(listener net.Listener){
	for  {
		conn, _ := listener.Accept()
		go ServerHandle(conn)
		sockets = append(sockets, SocketManager(conn))
	}
}
func ServerSendHandle(){
	var buf string
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		buf = stdin.Text()
		if buf == "exit"{
			break
		}
		SendAll(buf)
		fmt.Println(buf)
	}
}