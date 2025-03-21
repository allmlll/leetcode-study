package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*node
	tail, head *node
}

type node struct {
	next, prev *node
	key, value int
}

func initNode(key, value int) *node {
	return &node{
		key:   key,
		value: value,
	}
}
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
		tail:     initNode(0, 0),
		head:     initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) addToHead(n *node) {
	this.head.next.prev = n
	n.next = this.head.next
	this.head.next = n
	n.prev = this.head
}
func (this *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
func (this *LRUCache) removeToHead(n *node) {
	this.remove(n)
	this.addToHead(n)

}
func (this *LRUCache) removeTail() *node {
	n := this.tail.prev
	this.remove(n)
	return n
}
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	n := this.cache[key]
	this.removeToHead(n)
	return n.value
}
func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok {
		n := initNode(key, value)
		this.cache[key] = n
		this.addToHead(n)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.key)
			this.size--
		}

	} else {
		n := this.cache[key]
		n.value = value
		this.removeToHead(n)
	}
}
