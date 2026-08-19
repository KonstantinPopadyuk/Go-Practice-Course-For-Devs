package task14

import "time"

type Cache[K comparable, V any] struct {
	// TODO: add synchronization and storage fields.
	now func() time.Time
}

func New[K comparable, V any](now func() time.Time) *Cache[K, V] {
	// TODO: implement. If now is nil, use time.Now.
	return &Cache[K, V]{now: now}
}

func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
	// TODO: implement.
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	// TODO: implement.
	var zero V
	return zero, false
}

func (c *Cache[K, V]) Delete(key K) {
	// TODO: implement.
}

func (c *Cache[K, V]) Len() int {
	// TODO: implement.
	return 0
}
