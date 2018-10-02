package vbcore

import (
	"testing"
	"time"
)

func TestStringCacheAdd(t *testing.T) {
	refreshFunc := func() string {
		return "value"
	}

	c := NewStringCache()
	c.Add("key1", 1, refreshFunc)
	c.Add("key2", 1, refreshFunc)
	c.Add("key3", 1, refreshFunc)

	if _, ok := c.storage["key1"]; !ok {
		t.Fail()
	}
	if _, ok := c.storage["key2"]; !ok {
		t.Fail()
	}
	if _, ok := c.storage["key3"]; !ok {
		t.Fail()
	}
}

func TestStringCacheGet(t *testing.T) {
	c := NewStringCache()
	c.Add("key1", 1, func() string {
		return "value1"
	})
	c.Add("key2", 1, func() string {
		return "value2"
	})
	c.Add("key3", 1, func() string {
		return "value3"
	})

	c.Get("key1")
	c.Get("key2")
	c.Get("key3")

	time.Sleep(time.Duration(100) * time.Millisecond)

	if c.Get("key1") != "value1" {
		t.Fail()
	}
	if c.Get("key2") != "value2" {
		t.Fail()
	}
	if c.Get("key3") != "value3" {
		t.Fail()
	}
}

func TestStringCacheBackgroundSync(t *testing.T) {
	c := NewStringCache()
	c.Add("something", 1, func() string {
		return "value"
	})

	val := c.Get("something")
	if val != "" {
		t.Fail()
	}

	time.Sleep(time.Duration(100) * time.Millisecond)

	val = c.Get("something")
	if val != "value" {
		t.Fail()
	}
}
