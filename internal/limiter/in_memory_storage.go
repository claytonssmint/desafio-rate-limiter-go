package limiter

import (
	"sync"
	"time"
)

type InMemoryStorage struct {
	data  map[string]int
	mutex sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]int),
	}
}

func (s *InMemoryStorage) Increment(key string) (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[key]++
	return s.data[key], nil
}

func (s *InMemoryStorage) SetExpiration(key string, ttl time.Duration) error {
	// Simulação simples sem expiração real
	return nil
}

func (s *InMemoryStorage) GetCount(key string) (int, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.data[key], nil
}

func (s *InMemoryStorage) Reset(key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.data, key)
	return nil
}
