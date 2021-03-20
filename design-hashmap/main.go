package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(770, 2)
	param_2 := obj.Get(1)
	obj.Remove(770)
	fmt.Println(param_2)
}

type Entry struct {
	Key int
	Val int
}

const buckets_length = 769

type MyHashMap struct {
	Buckets []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{Buckets: make([]list.List, buckets_length)}
}

func (this *MyHashMap) Hash(key int) int {
	return key % buckets_length
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := this.Hash(key)
	l := this.Buckets[index]
	for e := l.Front(); e != nil; e = e.Next() {
		if entry := e.Value.(*Entry); entry.Key == key {
			entry.Val = value
			return
		}
	}
	this.Buckets[index].PushBack(&Entry{Key: key, Val: value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := this.Hash(key)
	l := this.Buckets[index]
	for e := l.Front(); e != nil; e = e.Next() {
		if entry := e.Value.(*Entry); entry.Key == key {
			return entry.Val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := this.Hash(key)
	l := this.Buckets[index]
	for e := l.Front(); e != nil; e = e.Next() {
		if entry := e.Value.(*Entry); entry.Key == key {
			l.Remove(e)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
