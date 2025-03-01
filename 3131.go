package main

func findMin(arr []int) int {
	if len(arr) == 0 {
		panic("Array is empty")
	}

	min := arr[0] // 初始化 min 为数组的第一个元素
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}
func addedInteger(nums1 []int, nums2 []int) int {
	return findMin(nums2) - findMin(nums1)
}
