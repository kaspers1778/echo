package main

import (
	"fmt"
	"log"
	"net"
)

const (
	TYPE = "tcp"
	PORT = "7"
	HOST = "localhost"
)

func main() {
	adr := fmt.Sprintf("%v:%v", HOST, PORT)
	l, err := net.Listen(TYPE, adr)
	if err != nil {
		log.Fatal("Error on listening: ", err)

	}
	log.Printf("Server is running on %v\n", adr)
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal("Error on connection: ", err)
		}
		go handleEcho(c)
	}
}

func handleEcho(c net.Conn) {
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		if _, err := c.Read(buf); err != nil {
			log.Println("Error on reading: ", err)
			break
		}
		log.Printf("Recive: %s", buf)
		if _, err := c.Write(buf); err != nil {
			log.Println("Error on writing: ", err)
			break
		}
	}
}
