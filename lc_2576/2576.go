package lc_2576

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	count := 0
	t := 25
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]/2 >= nums[0] {
			for j := t - 1; j >= 0; j-- {
				if nums[j] <= nums[i]/2 && nums[j] != -1 {
					nums[j] = -1
					nums[i] = -1
					count += 2
					t = j
					break
				}
			}
		}

	}
	return count
}

//func main() {
//	nums := []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}
//	maxNumOfMarkedIndices(nums)
//}
