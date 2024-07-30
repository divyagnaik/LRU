package main

import (
	"container/list"
	"fmt"
)

//use a combination of a doubly linked list and a hash map.
//The doubly linked list will help keep track of the order of access,
//the hash map will provide efficient access to the cache items.

// Cache structure with fixed capacity
type Cache struct {
	capacity int
	cache    map[string]*list.Element
	order    *list.List
}

// Entry is the type stored in the list
type Entry struct {
	key   string
	value interface{}
}

// NewCache initializes a new Cache with a given capacity
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

// Get retrieves the value for a key from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem)
		return elem.Value.(*Entry).value, true
	}
	return nil, false
}

// Put adds a key-value pair to the cache
func (c *Cache) Put(key string, value interface{}) {
	if elem, found := c.cache[key]; found {
		// Update the value and move to front
		entry := elem.Value.(*Entry)
		entry.value = value
		c.order.MoveToFront(elem)
		return
	}

	if c.order.Len() >= c.capacity {
		// Remove the oldest item
		oldest := c.order.Back()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.cache, oldest.Value.(*Entry).key)
		}
	}

	// Add new item
	entry := &Entry{key: key, value: value}
	elem := c.order.PushFront(entry)
	c.cache[key] = elem
}

func main() {
	c := NewCache(2)
	c.Put("1", "divya")

	for key, value := range c.cache {
		fmt.Println(key, value.Value)
	}

	fmt.Println()
	c.Put("2", "pallu")

	for key, value := range c.cache {
		fmt.Println(key, value.Value)
	}

	fmt.Println()
	c.Get("1")

	for key, value := range c.cache {
		fmt.Println(key, value.Value)
	}

	fmt.Println()
	c.Put("3", "akshu")

	for key, value := range c.cache {
		fmt.Println(key, value.Value)
	}

	fmt.Println()
	c.Put("4", "akshatha")

	for key, value := range c.cache {
		fmt.Println(key, value.Value)
	}
}
