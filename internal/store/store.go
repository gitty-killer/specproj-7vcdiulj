package store

import "sync"

type Store struct {
    mu sync.Mutex
    items map[string]string
}

func New() *Store {
    return &Store{items: map[string]string{}}
}

func (s *Store) Set(key, value string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.items[key] = value
}

func (s *Store) Get(key string) (string, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    v, ok := s.items[key]
    return v, ok
}
