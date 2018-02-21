package chat

import (
	"bufio"
	"os"
	"fmt"
)

func ReadHandle(socket *Socket){
	for{
		message := socket.Read()
		if message == ""{
			break
		}
		println("[受信]: "+ message)
	}
}

func SendHandle(socket *Socket){
	var buf string
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		buf = stdin.Text()
		if buf == "exit"{
			socket.Close()
			break
		}
		socket.Send(buf)
	}
}