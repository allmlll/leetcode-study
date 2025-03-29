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
func main() {
	list1 := arrayToLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	list2 := arrayToLinkedList([]int{9, 9, 9, 9})
	addTwoNumbers(list1, list2)
}
