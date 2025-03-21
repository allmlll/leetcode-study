package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < l-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return dummy.Next
}
