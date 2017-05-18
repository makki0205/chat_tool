package Chat

import (
	"bufio"
	"os"
	"fmt"
)

func ReadHandle(socket *Socket){
	for{
		message := socket.Read()
		println("[受信]: "+ message)
	}
}

func SendHandle(socket *Socket){
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		socket.Send(stdin.Text())
		fmt.Println("[送信]: " + stdin.Text())
	}
}