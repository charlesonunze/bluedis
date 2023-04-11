package main

import "sync"

type Store struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewStore() *Store {
	return &Store{
		mu:    sync.RWMutex{},
		store: map[string]string{},
	}
}

func (s *Store) Set(k, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[k] = v
}

func (s *Store) Get(k string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.store[k]
}
