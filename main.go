package main

func listinit(arr []int) *ListNode {
	var head *ListNode
	p := head
	for _, v := range arr {

		p = &ListNode{
			Val:  v,
			Next: &ListNode{},
		}
		p = p.Next
	}
	return head
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	// 创建头节点
	head := &ListNode{Val: arr[0]}
	current := head

	// 遍历数组，创建链表节点并连接
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i]}
		current.Next = node
		current = node
	}

	return head
}
func ArrayToTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil // 处理空数组或根节点为空的情况
	}

	// 创建根节点
	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root} // 用队列辅助层序构建
	i := 1                     // 当前处理的数组索引

	// 循环处理队列中的每个节点
	for len(queue) > 0 && i < len(arr) {
		node := queue[0]  // 取出队首节点
		queue = queue[1:] // 更新队列

		// 构建左子节点（如果数组未越界且值不为-1）
		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++ // 索引后移，无论是否创建节点

		// 构建右子节点（如果数组未越界且值不为-1）
		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++ // 索引后移，无论是否创建节点
	}

	return root
}
func main() {
	root := ArrayToTree([]int{3, 9, 20, -1, -1, 15, 7})
	maxDepth(root)
}
