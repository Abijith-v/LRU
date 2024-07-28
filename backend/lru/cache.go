package lru

import (
    "container/list"
    "sync"
    "time"
)

type CacheItem struct {
    Key       string
    Value     string
    Timestamp time.Time
}

type LRUCache struct {
    capacity int
    ttl      time.Duration
    cache    map[string]*list.Element
    list     *list.List
    mutex    sync.Mutex
}

func CreateLRUCache(capacity int, ttl time.Duration) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        ttl:      ttl,
        cache:    make(map[string]*list.Element),
        list:     list.New(),
    }
}

func (c *LRUCache) Get(key string) (string, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if element, ok := c.cache[key]; ok {
        item := element.Value.(*CacheItem)
        if time.Since(item.Timestamp) > c.ttl {
            // Item has expired
            c.list.Remove(element)
            delete(c.cache, key)
            return "", false
        }

        c.list.MoveToFront(element)
        return item.Value, true
    }
    return "", false
}

func (c *LRUCache) Put(key, value string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if element, ok := c.cache[key]; ok {
        c.list.MoveToFront(element)
        item := element.Value.(*CacheItem)
        item.Value = value
        item.Timestamp = time.Now()
        return
    }

    if c.list.Len() >= c.capacity {
        evict := c.list.Back()
        if evict != nil {
            c.list.Remove(evict)
            delete(c.cache, evict.Value.(*CacheItem).Key)
        }
    }

    item := &CacheItem{Key: key, Value: value, Timestamp: time.Now()}
    element := c.list.PushFront(item)
    c.cache[key] = element
}

// deletes the least recently used key
func (c *LRUCache) DeleteLRU() {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    evict := c.list.Back()
    if evict != nil {
        c.list.Remove(evict)
        delete(c.cache, evict.Value.(*CacheItem).Key)
    }
}

// Keys method to return all keys in the cache
func (c *LRUCache) Keys() []string {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    var keys []string
    for element := c.list.Front(); element != nil; element = element.Next() {
        keys = append(keys, element.Value.(*CacheItem).Key)
    }
    return keys
}
