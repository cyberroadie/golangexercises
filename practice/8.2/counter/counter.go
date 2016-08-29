package main

import (
	"io"
	"log"
	"net"
	"strconv"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:1000")
	if err != nil {
		log.Fatalf("%s", err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("%s", err)
		}
		doSomething(c)
	}
}

func doSomething(c net.Conn) {
	defer c.Close()
	count := 0
	for {
		_, err := io.WriteString(c, strconv.Itoa(count)+"\n")
		count++
		if err != nil {
			return
		}
	}
}
