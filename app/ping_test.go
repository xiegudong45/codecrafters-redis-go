package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestMain(t *testing.T) {
	client := NewRedisClient()
	for i := 1; i <= 3; i++ {
		if err := runPing(client, 1); err != nil {

			fmt.Printf("error is: %+v\n", err)
		}
	}

}

func runPing(client *redis.Client, clientNum int64) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}

	if pong != "PONG" {
		return fmt.Errorf("client-%d: Expected \"PONG\", got %#v", clientNum, pong)
	}

	return nil
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		DialTimeout: 5 * time.Second,
		Dialer: func() (net.Conn, error) {
			attempts := 0

			for {
				var err error
				var conn net.Conn

				conn, err = net.Dial("tcp", "localhost:6379")

				if err == nil {
					return conn, nil
				}

				// Already a timeout
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					return nil, err
				}

				// 50 * 100ms = 5s
				if attempts > 50 {
					return nil, err
				}

				attempts += 1
				time.Sleep(100 * time.Millisecond)
			}
		},
	})
}
