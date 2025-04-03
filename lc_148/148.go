package lc_148

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	l := getListLength(head)
	if l == 0 {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	for splitLength := 1; splitLength < l; splitLength <<= 1 {
		prev := dummy
		cur := dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; cur.Next != nil && i < splitLength; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; cur.Next != nil && i < splitLength; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = mergeList(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next

		}

	}
	return dummy.Next
}
func getListLength(head *ListNode) (len int) {
	for p := head; p != nil; p = p.Next {
		len++
	}
	return
}
func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	temp := newHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}
	if list1 == nil {
		temp.Next = list2
	}
	if list2 == nil {
		temp.Next = list1
	}
	return newHead.Next
}
