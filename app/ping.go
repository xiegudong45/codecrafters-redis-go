package main

import (
	"bufio"
	"io"
	"net"
	"strings"
)

func HandlePing(c net.Conn) {
	defer c.Close()
	for {
		reader := bufio.NewReader(c)
		for {
			// read command from terminal
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				break
			}
			// make response
			if strings.Contains(line, "ping") {
				c.Write([]byte("+PONG\r\n"))
			}
		}
	}
}
