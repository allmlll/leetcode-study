package main

import "slices"

func rotate(nums []int, k int) {
	k %= len(nums) // 轮转 k 次等于轮转 k % n 次
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
