package lc_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}

	for i := 1; i < l; i <<= 1 {
		for p := 0; p+i < l; p += 2 * i {
			lists[p] = MergeList(lists[p], lists[p+i])
		}

	}
	return lists[0]
}
func MergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	if list1 == nil {
		p.Next = list2
	} else if list2 == nil {
		p.Next = list1
	}
	return dummy.Next
}
