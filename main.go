package main

import (
	"os"
	"fmt"
	"github.com/makki0205/chat_tool/chat"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}
	if os.Args[1] == "-l"{
		chat.ServerInit()
	}else{
		chat.ClientInit(os.Args[1])
	}
}