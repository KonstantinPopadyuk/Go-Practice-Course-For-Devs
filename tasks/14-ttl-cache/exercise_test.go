package task14

import (
	"sync"
	"testing"
	"time"
)

func TestCacheExpiry(t *testing.T) {
	current := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	c := New[string, int](func() time.Time { return current })
	c.Set("answer", 42, time.Minute)
	if got, ok := c.Get("answer"); !ok || got != 42 {
		t.Fatalf("Get = (%d, %v)", got, ok)
	}
	current = current.Add(time.Minute)
	if _, ok := c.Get("answer"); ok || c.Len() != 0 {
		t.Fatal("expired item is still present")
	}
	c.Set("dead", 1, 0)
	if _, ok := c.Get("dead"); ok {
		t.Fatal("non-positive TTL item is live")
	}
}

func TestCacheLenSweepsExpiredEntries(t *testing.T) {
	current := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	c := New[string, int](func() time.Time { return current })
	c.Set("expired", 1, time.Second)
	c.Set("live", 2, time.Hour)
	current = current.Add(time.Second)
	if got := c.Len(); got != 1 {
		t.Fatalf("Len = %d, want only the live item", got)
	}
	c.Delete("live")
	if _, ok := c.Get("live"); ok || c.Len() != 0 {
		t.Fatal("Delete did not remove live item")
	}
}

func TestCacheExpiredSetReplacesLiveValue(t *testing.T) {
	c := New[string, int](func() time.Time { return time.Unix(100, 0) })
	c.Set("key", 1, time.Hour)
	c.Set("key", 2, 0)
	if value, ok := c.Get("key"); ok || c.Len() != 0 {
		t.Fatalf("Get after expired replacement = (%d, %v), Len = %d", value, ok, c.Len())
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	c := New[int, int](nil)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			c.Set(n%4, n, time.Second)
			_, _ = c.Get(n % 4)
			if n%3 == 0 {
				c.Delete(n % 4)
			}
		}(i)
	}
	wg.Wait()
	_ = c.Len()
}
