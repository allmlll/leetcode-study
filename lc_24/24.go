package lc_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	c := dummy
	for c.Next != nil && c.Next.Next != nil {
		pre := c.Next
		nxt := pre.Next
		pre.Next = nxt.Next
		nxt.Next = pre
		c.Next = nxt
		c = pre
	}
	return dummy.Next

}
