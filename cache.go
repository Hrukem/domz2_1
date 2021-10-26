package domz2_1

import (
	"sync"
	"time"
)

type data struct {
	value    interface{}
	lifetime int64
}

type Cache struct {
	c  map[string]*data
	mu sync.Mutex
}

// New function create new cache and start goroutine for check
// lifetime of the data
func New() *Cache {
	res := &Cache{c: make(map[string]*data)}
	go res.checkTime()
	return res
}

// Set function implementation add data in cache
func (m *Cache) Set(key string, val interface{}, ttl int) {
	m.mu.Lock()
	it := &data{value: val}
	m.c[key] = it
	m.c[key].lifetime =
		time.Now().Add(time.Second * time.Duration(ttl)).Unix()
	m.mu.Unlock()
}

// Get function implementation get data from cache
func (m *Cache) Get(key string) interface{} {
	val, ok := m.c[key]
	if !ok {
		return nil
	}
	return val.value
}

// Delete function implementation delete data from cache
func (m *Cache) Delete(key string) {
	m.mu.Lock()
	_, ok := m.c[key]
	if ok {
		delete(m.c, key)
	}
	m.mu.Unlock()
}

// checkTime function checks the lifetime of the data
func (m *Cache) checkTime() {
	for {
		time.Sleep(time.Millisecond * 500)
		for k, v := range m.c {
			if v.lifetime < time.Now().Unix() {
				m.Delete(k)
			}
		}
	}
}
