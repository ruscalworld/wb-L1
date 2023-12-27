package task7

import "sync"

// Наиболее простое решение для этой задачи - sync.Map
// Но не будем искать лёгкие пути...

type SyncMap[K comparable, V any] struct {
	lock *sync.RWMutex
	m    map[K]V
}

func MakeMap[K comparable, V any](size int) *SyncMap[K, V] {
	return &SyncMap[K, V]{
		lock: &sync.RWMutex{},
		m:    make(map[K]V, size),
	}
}

func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	v, ok := m.m[key]
	return v, ok
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.m[key] = value
}
