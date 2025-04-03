package lc_238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	L := make([]int, n)
	R := make([]int, n)
	res := make([]int, n)
	L[0] = 1
	R[n-1] = 1
	for i := 1; i < n; i++ {
		L[i] = nums[i-1] * L[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}
	for i := 0; i < n; i++ {
		res[i] = L[i] * R[i]
	}
	return res
}
