package mutex

import "sync"

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Neuen Wert setzen (Schreibsperre)
func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Wert abfragen (Lesesperre)
func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}
