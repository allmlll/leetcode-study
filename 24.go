package main

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre, nxt := dummy.Next, dummy.Next.Next
	for pre != nil && nxt != nil {
		if nxt.Next == nil {
			nxt.Next = pre
			pre.Next = nil
		} else {
			pre.Next = nxt.Next
			nxt.Next = pre
			pre = pre.Next
			nxt = pre.Next
		}

	}
	return dummy.Next

}
