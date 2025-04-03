package lc_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	c := &ListNode{}
	dummy := c
	t := 0
	for p1 != nil || p2 != nil {
		if p1 == nil {
			if p2.Val+t >= 10 {
				tList := &ListNode{
					Val: p2.Val + t - 10,
				}
				dummy.Next = tList
				dummy = tList
				t = 1
				p2 = p2.Next
			} else {
				tList := &ListNode{
					Val: p2.Val + t,
				}
				dummy.Next = tList
				dummy = tList
				t = 0
				p2 = p2.Next
			}

		} else if p2 == nil {
			if p1.Val+t >= 10 {
				tList := &ListNode{
					Val: p1.Val + t - 10,
				}
				dummy.Next = tList
				dummy = tList
				t = 1
				p1 = p1.Next
			} else {
				tList := &ListNode{
					Val: p1.Val + t,
				}
				dummy.Next = tList
				dummy = tList
				t = 0
				p1 = p1.Next
			}
		} else {

			sum := p1.Val + p2.Val
			if sum+t >= 10 {
				tList := &ListNode{
					Val: sum - 10 + t,
				}
				dummy.Next = tList
				dummy = tList
				t = 1
			} else {
				tList := &ListNode{
					Val: sum + t,
				}
				dummy.Next = tList
				dummy = tList
				t = 0
			}
			p1 = p1.Next
			p2 = p2.Next

		}

	}
	if t != 0 {
		tList := &ListNode{
			Val: t,
		}
		dummy.Next = tList
		dummy = tList
	}
	return c.Next

}
