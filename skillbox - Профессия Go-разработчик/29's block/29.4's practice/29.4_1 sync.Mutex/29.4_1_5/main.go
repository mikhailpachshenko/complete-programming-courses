package main

import (
	"sync"
)

type server struct {
	status string
	mu     sync.Mutex
}

func main() {
	s := server{}
	for i := 0; i < 1000; i++ {
		go s.Alive()
		go s.Down()
	}
}

func (s *server) Alive() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = "Alive"
}

func (s *server) Down() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = "Down"
}
