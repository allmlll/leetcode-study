package main

type BookMyShow struct {
	n    int
	m    int
	Show []int
	sum  int
}

//func Constructor(n int, m int) BookMyShow {
//
//	var show BookMyShow
//	ints := make([]int, n)
//	for i := 0; i < n; i++ {
//		ints[i] = m
//		show.sum += m
//	}
//	show.n = n
//	show.m = m
//	show.Show = ints
//	return show
//}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	if k > this.m {
		return []int{}
	}
	for i := 0; i <= maxRow; i++ {
		if this.Show[i] >= k {
			this.Show[i] = this.Show[i] - k
			return []int{i, this.m - this.Show[i] - k}
		}
	}
	return []int{}

}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if this.sum < k {
		return false
	}
	var sum int
	for i := 0; i <= maxRow; i++ {
		sum += this.Show[i]
	}
	if k > sum {
		return false
	} else {
		this.sum -= k
		for i := 0; i <= maxRow; i++ {
			if this.Show[i] >= 0 {
				if this.Show[i]-k <= 0 {
					k = k - this.Show[i]
					this.Show[i] = 0

				} else {
					this.Show[i] = this.Show[i] - k
					k = 0
				}
			}
			if k <= 0 {
				return true
			}
		}
	}
	return false
}

//func main() {
//	constructor := Constructor(30, 583)
//	constructor.Scatter(699, 10)
//	constructor.Gather(507, 24)
//	constructor.Scatter(151, 15)
//	constructor.Gather(753, 19)
//	constructor.Gather(384, 4)
//	constructor.Gather(178, 1)
//	constructor.Gather(962, 3)
//
//}
