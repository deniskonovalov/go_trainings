package main

import (
	"fmt"
	"lesson_08/lru"
)

func main() {
	c := lru.NewLruCache(3)

	c.Put("key", "value")

	c.Put("key2", "value2")
	c.Put("key4", "value11111")
	c.Put("key5", "value11111")
	c.Put("key6", "value11111")
	c.Put("key7", "value11111")

	c.Put("key3", "value11111")

	fmt.Println(c.Get("key7"))
	fmt.Println(c.Get("key"))
}
