package lc_138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cachedNode = make(map[*Node]*Node)
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cachedNode[node]; ok {
		return n
	}
	n := &Node{
		Val: node.Val,
	}
	cachedNode[node] = n
	n.Next = deepCopy(node.Next)
	n.Random = deepCopy(node.Random)
	return n
}
