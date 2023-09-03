package go_catch

import (
	"fmt"
	"sync"
	"time"
)

// Cache implements a simple in-memory cache used token check
// expire at least 60 seconds for cpu load
type Cache struct {
	sync.Mutex
	mapping       map[string]*Value
	expire        int
	end           chan int
	checkInterval int
	startTime     int64
}

// Value cache value which created time
type Value struct {
	value   string
	created int64
}

// Init cache,
func (c *Cache) Init() {
	c.startTime = time.Now().Unix()
	c.mapping = make(map[string]*Value)
	c.end = make(chan int, 1)
	if c.checkInterval == 0 {
		c.checkInterval = 60
	}
	if c.expire == 0 {
		c.expire = 60
	}
	go c.checkExpired()
}

func (c *Cache) checkExpired() {
	for {
		select {
		case <-time.After(time.Duration(c.checkInterval) * time.Second):
			c.Lock()
			now := time.Now().Unix()
			var cleankeys []string
			for k, v := range c.mapping {
				if v != nil {
					if now-v.created > int64(c.expire) {
						cleankeys = append(
							cleankeys, k)
					}
				}
			}
			for i := 0; i < len(cleankeys); i++ {
				_, ok := c.mapping[cleankeys[i]]
				if ok {
					delete(c.mapping, cleankeys[i])
				}
			}
			c.Unlock()
		case <-c.end:
			return
		}
	}
}

// Close end goroutine
func (c *Cache) Close() {
	c.end <- 1
}

// Get get cache
func (c *Cache) Get(key string) string {
	c.Lock()
	defer c.Unlock()
	if c.mapping == nil {
		fmt.Println("cache not init")
		return ""
	}
	v, ok := c.mapping[key]
	if ok {
		return v.value
	}
	return ""
}

// Set set cache
func (c *Cache) Set(key string, val string) {
	c.Lock()
	defer c.Unlock()
	if c.mapping == nil {
		fmt.Println("cache not init")
		return
	}
	v := Value{val, time.Now().Unix()}
	c.mapping[key] = &v
}

// Delete delete cache
func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	_, ok := c.mapping[key]
	if ok {
		delete(c.mapping, key)
	}
}
