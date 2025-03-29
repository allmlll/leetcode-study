package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head.Next,
	}
	prev := dummy

	for {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nex := tail.Next
		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = nex
		prev = tail
		head = nex
	}
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	nex := tail.Next
	prev := nex
	p := head
	for prev != tail {
		nax := p.Next
		p.Next = prev
		prev = p
		p = nax
	}
	return tail, head
}
