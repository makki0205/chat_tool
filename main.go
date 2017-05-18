package main

import (
	"os"
	"./Chat"
	"fmt"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}
	if os.Args[1] == "-l"{
		Chat.ServerInit()
	}else{
		Chat.ClientInit(os.Args[1])
	}
}