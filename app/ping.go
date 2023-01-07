package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
)

//	func HandlePing(c net.Conn, store *Store) {
//		defer c.Close
func HandlePing(c net.Conn, store *Store) {
	defer c.Close()

	for {
		value, err := DecodeRESP(bufio.NewReader(c))
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("Error handling RESP", err.Error())
			return
		}
		command := value.Array()[0].String()
		args := value.Array()[1:]
		fmt.Println("args:", args)
		fmt.Println("command: ", command)

		switch command {
		case "ping":
			if len(args) == 0 {
				c.Write([]byte("+PONG\r\n"))
			} else {
				c.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0].String()), args[0].String())))
			}
		case "echo":
			c.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0].String()), args[0].String())))
		case "set":
			store.Set(args[0].String(), args[1].String())
			c.Write([]byte("+OK\r\n"))
		case "get":
			c.Write([]byte(fmt.Sprintf("+%s\r\n", store.Get(args[0].String()))))
		default:
			c.Write([]byte("-ERR unknown command '" + command + "'\r\n"))
		}
	}

}
