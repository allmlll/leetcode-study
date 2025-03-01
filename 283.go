package main

func moveZeroes(nums []int) {
	i, j, n := 0, 0, len(nums)-1
	for j < n {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[j]
			i++
		}
		j++
	}
}
