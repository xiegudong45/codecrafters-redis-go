package main

import "fmt"

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) string {
	res := fmt.Sprintf("%s\r\n", s.data[key])
	return res
}

func (s *Store) Set(key string, value string) string {
	s.data[key] = value
	return "+OK\r\n"
}
