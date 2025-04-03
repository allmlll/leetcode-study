package lc_2708

import (
	"fmt"
	"sort"
)

func maxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}

	count := 0
	res := 1
	flag := 1
	i := 0
	k := 0
	for _, num := range nums {
		if num < 0 {
			count++
		}
	}
	if count%2 == 1 {
		count--
	}
	m := count
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, num := range nums {
		if num > 0 {
			res = res * num
			k++
		}
		if num < 0 {
			if count == 0 {
				continue
			}
			res = res * nums[len(nums)-1-i]
			i++
			count--
		}
		if num == 0 {
			flag = 0
		}

	}
	if k == 0 && flag == 0 && m == 0 {
		res = 0
	}

	fmt.Println(res)
	return int64(res)

}

//func main() {
//	nums := []int{0, -1}
//	maxStrength(nums)
//}
