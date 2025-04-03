package lc_338

func countSetBit(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1 // 将 n 的最右边的一个 1 变为 0
		count++
	}
	return count
}
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res = append(res, countSetBit(i))
	}
	return res
}
