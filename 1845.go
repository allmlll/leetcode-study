package main

type SeatManager struct {
	Root *TreeNodea
}
type TreeNodea struct {
	Value  int
	Status bool
	Left   *TreeNodea
	Right  *TreeNodea
}

// Constructor 构造函数，用于初始化 SeatManager
func Constructor(n int) SeatManager {
	return SeatManager{
		Root: buildBST(1, n),
	}
}
func buildBST(start, end int) *TreeNodea {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNodea{Value: mid, Status: false}
	root.Left = buildBST(start, mid-1)
	root.Right = buildBST(mid+1, end)
	return root
}

// AllocateSeat 分配一个可用座位
//func (this *SeatManager) Reserve() int {
//	var root = this.Root
//
//	// 如果没有找到符合条件的节点，则返回 nil
//	return 0
//}

// FreeSeat 释放一个已分配的座位
func (this *SeatManager) Unreserve(seatNumber int) {
	var root = this.Root
	for root != nil {
		if root.Value == seatNumber {
			root.Status = false
			break
		} else if seatNumber < root.Value {
			root = root.Left
		} else {
			root = root.Right
		}
	}
}

//func main() {
//	constructor := Constructor(5)
//	fmt.Println(constructor.Reserve())
//	constructor.Reserve()
//	constructor.Unreserve(2)
//}
