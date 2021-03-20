package main

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4
}

type DLinkedNode struct {
	Key, Val   int
	prev, next *DLinkedNode
}

func NewDLinkedNode(key int, val int) *DLinkedNode {
	return &DLinkedNode{Key: key, Val: val}
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     NewDLinkedNode(0, 0),
		tail:     NewDLinkedNode(0, 0),
	}

	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head

	return lruCache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.MoveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.Val = value
		this.MoveToHead(node)
		return
	}
	node = NewDLinkedNode(key, value)
	this.cache[key] = node
	this.AddNodeToHead(node)
	this.size++
	if this.size > this.capacity {
		removedNode := this.RemoveTail()
		delete(this.cache, removedNode.Key)
		this.size--
	}
}

func (this *LRUCache) AddNodeToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) RemoveNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) MoveToHead(node *DLinkedNode) {
	this.RemoveNode(node)
	this.AddNodeToHead(node)
}

func (this *LRUCache) RemoveTail() *DLinkedNode {
	node := this.tail.prev
	this.RemoveNode(node)
	return node
}
