package main

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		nex := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next

	}
	return dummy.Next
}

func reverse1(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nex := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next

	}
	return dummy.Next
}
