package main

import (
	"sort"
)

func threeSum(nums []int) (res [][]int) {
	first := 0
	second := 1
	third := len(nums) - 1
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})

	for {
		second = first + 1
		for {
			third = len(nums) - 1
			for {
				if second >= third {
					break
				}
				sum := nums[first] + nums[second]
				if sum+nums[third] > 0 {
					third--
				}
				if sum+nums[third] == 0 {
					res = append(res, []int{nums[first], nums[second], nums[third]})
					break
				}
				if sum+nums[third] < 0 {
					break
				}

			}
			i := 1
			for nums[second+i] == nums[second] {
				i++

			}
			second = second + i

			if second > len(nums)-2 {
				break
			}
		}

		j := 1
		for nums[first+j] == nums[first] {
			j++

		}
		first = first + j

		if first > len(nums)-3 {
			break
		}

	}
	return res
}
