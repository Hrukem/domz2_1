package domz2_1

import "sync"

type Cache struct {
	cacheMap map[string]interface{}
	mu       sync.Mutex
}

func NewCache() Cache {
	return Cache{cacheMap: map[string]interface{}{}}
}

// Set function implementation add data in cache
func (m *Cache) Set(key string, value interface{}) {
	m.mu.Lock()
	m.cacheMap[key] = value
	m.mu.Unlock()
}

// Get function implementation get data from cache
func (m *Cache) Get(key string) interface{} {
	return m.cacheMap[key]
}

// Delete function implementation delete data from cache
func (m *Cache) Delete(key string) {
	m.mu.Lock()
	delete(m.cacheMap, key)
	m.mu.Unlock()
}
