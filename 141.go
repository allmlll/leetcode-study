package main

func hasCycle(head *ListNode) bool {
	slow, fast := head, head.Next
	for slow != fast {
		if fast.Next == nil || fast == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
