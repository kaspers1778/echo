package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	TYPE = "tcp"
	PORT = "7"
	HOST = "localhost"
)

func main() {
	adr := fmt.Sprintf("%v:%v", HOST, PORT)
	c, err := net.Dial(TYPE, adr)
	if err != nil {
		log.Fatal("Error on connecting: ", err)
	}

	for {
		br := bufio.NewReader(os.Stdin)
		fmt.Print("Echo: ")
		message, _ := br.ReadString('\n')
		fmt.Fprintf(c, message+"\n")
		echo, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print(echo)
	}
}
