# 14. A Concurrency-Safe TTL Cache

**Time:** 35–40 minutes.

Implement a generic `Cache[K, V]` whose entries have a time to live.

## API and Contract

- `New(now func() time.Time)` creates a cache; injecting the clock makes tests deterministic.
- `Set(key, value, ttl)` stores a value; `ttl <= 0` means the entry is already expired.
- `Get(key) (V, bool)` lazily removes an expired entry.
- `Delete(key)` removes a key.
- `Len()` counts only live entries and removes expired ones.
- All methods must be safe for concurrent use.
- Do not hold a lock longer than necessary, and never copy a mutex after it has been used.

Run: `go test ./tasks/14-ttl-cache -race -v`.

For simplicity, no background cleanup goroutine is required. Choose either `sync.Mutex` or `sync.RWMutex`, but remember that `Get` sometimes removes an entry and is therefore not always a read-only operation.

## Going Further

Add a size limit and LRU eviction; test the access order separately.
