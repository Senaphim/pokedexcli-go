package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com/test",
			value: []byte("somedata"),
		},
		{
			key:   "https://example.com/test2",
			value: []byte("somemoredata"),
		},
	}

	for _, c := range cases {
		cache := NewCache(interval)
		cache.Add(c.key, c.value)
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("Expected to find value for key: %s", c.key)
		}
		if string(val) != string(c.value) {
			t.Errorf("Value incorrect for key: %s. Expected: %s, got: %s",
				c.key, string(c.value), string(val))
		}
	}
	return
}

func TestReapMethod(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond
	cache := NewCache(baseTime)
	key := "https://example.com/test"
	cache.Add(key, []byte("somedata"))

	if _, ok := cache.Get(key); !ok {
		t.Errorf("Expected to find value while testing reap method")
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get(key); ok {
		t.Errorf("Expected not to find value while testing reap method")
	}
}
