package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandlePing(c net.Conn, store *Store) {
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
			} else if firstChar == "*3" {
				// store[]
				c.Write([]byte("OK"))
				break
			}
			idx++
		}
	}
}

func HandlePing1(c net.Conn, store *Store) {
	defer c.Close()
	buf := make([]byte, 1024)
	for {
		len, _ := c.Read(buf)
		command := string(buf[:len])
		wordsList := strings.Split(command, "\r\n")
		fmt.Println(wordsList)
		if wordsList[0] == "*1" && strings.ToUpper(wordsList[2]) == "PING" {
			c.Write([]byte("+PONG\r\n"))
			break
		} else if wordsList[0] == "*2" && strings.ToUpper(wordsList[2]) == "PING" || strings.ToUpper(wordsList[2]) == "ECHO" {
			res := fmt.Sprintf("%s\r\n%s\r\n", wordsList[3], wordsList[4])
			c.Write([]byte(res))
			break
		} else if wordsList[2] == "GET" {
			key := wordsList[4]
			c.Write([]byte(store.Get(key)))
			break
		} else if wordsList[2] == "SET" {
			key := wordsList[4]
			val := wordsList[6]
			c.Write([]byte(store.Set(key, val)))
			break
		}
	}

}
