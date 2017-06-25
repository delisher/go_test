package main

import (
	"bufio"
	"fmt"
	"net"
)

// main listen for new tcp connection on localhost 8053 port
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("Listen failure: %s\n", err.Error())
	}
	for {
		if c, err := l.Accept(); err == nil {
			go Echo(c)
		}
	}
}

// Echo handles connection
func Echo(c net.Conn) {
	defer c.Close()
	buf, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Read failure: %s | value: %+v\n", err.Error(), buf)
		return
	}
	fmt.Printf("Read value: %+v\n", buf)
	w, err := c.Write([]byte(buf))
	if err != nil {
		fmt.Printf("Write failure: %s\n", err.Error())
		return
	}
	fmt.Printf("Write value: %+v\n", w)
}
