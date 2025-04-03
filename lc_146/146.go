package lc_146

type Node struct {
	key, value int
	prev, next *Node
}
type LRUCache struct {
	head, tail     *Node
	cache          map[int]*Node
	size, capacity int
}

func initNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
		cache:    make(map[int]*Node),
		size:     0,
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	} else {
		n := this.cache[key]
		this.moveToHead(n)
		return n.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initNode(key, value)
		this.cache[key] = node
		this.moveToHead(node)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.key)
			this.size--
		}
	} else {

		n := this.cache[key]
		n.value = value
		this.moveToHead(n)

	}
}

func (this *LRUCache) addToHead(node *Node) {
	this.head.next.prev = node
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *Node {
	p := this.tail.prev
	this.remove(p)
	return p
}

func (this *LRUCache) moveToHead(node *Node) {
	this.addToHead(node)
	this.remove(node)
}
