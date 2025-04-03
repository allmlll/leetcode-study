package lc_234

type ListNode struct {
	Val  int
	Next *ListNode
}

func rreverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	NewHead := rreverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return NewHead
}
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for {
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		if fast.Next.Next == nil {
			fast = fast.Next
			break
		}
		fast = fast.Next.Next

	}
	newHead := rreverse(slow)
	p1 := head
	p2 := newHead
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
