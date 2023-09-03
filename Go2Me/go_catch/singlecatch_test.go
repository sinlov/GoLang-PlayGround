package go_catch

import (
	"fmt"
	"testing"
	"time"
)

func TestSingleCache(t *testing.T) {
	cache := Cache{expire: 1, checkInterval: 1}
	cache.Init()
	defer cache.Close()
	cache.Set("name", "myTest")
	v := cache.Get("name")
	if v != "myTest" {
		t.Fatal("cache get and set error")
	} else {
		fmt.Printf("get catch %s\n", v)
	}
	time.Sleep(3 * time.Second)
	v = cache.Get("name")
	if v != "" {
		t.Fatalf("cache set expire error %v", v)
	} else {
		t.Logf("error cache empty")
	}

}
