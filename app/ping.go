package main

import (
	"bufio"
	"fmt"
	"net"
)

func HandlePing(c net.Conn) {
	defer c.Close()
	for {
		scanner := bufio.NewScanner(c)
		var words []string
		idx := 0
		var firstChar string
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println(text)
			words = append(words, text)
			if idx == 0 {
				firstChar = text
			}
			if firstChar == "*1" && idx == 2 {
				c.Write([]byte("+PONG\r\n"))
				break
			} else if firstChar == "*2" && idx == 4 {
				res := fmt.Sprintf("%s\r\n%s\r\n", words[len(words)-2], words[len(words)-1])
				c.Write([]byte(res))
				break
			}
			idx++
		}
	}
}
