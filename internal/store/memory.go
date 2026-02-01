package store

import "sync"

type Store interface {
	Save(code string, url string) error
	Get(code string) (string, bool)
	Exists(code string) bool
}

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMemoryStore() *MemoryStore {

	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (m *MemoryStore) Save(code string, url string) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[code] = url
	return nil
}

func (m *MemoryStore) Get(code string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	url, ok := m.data[code]

	return url, ok
}

func (m *MemoryStore) Exists(code string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.data[code]

	return ok
}
