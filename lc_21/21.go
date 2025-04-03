package lc_21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := ListNode{}
	cur := &dummy
	for list1.Next != nil || list2.Next != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
		} else {
			cur.Next = list2
		}
		cur = cur.Next
	}
	if list1.Next != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next

}
