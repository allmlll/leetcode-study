package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	mp[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if j, ok := mp[target-nums[i]]; ok {
			return []int{j, i}
		}
		mp[nums[i]] = i
	}
	return []int{-1, -1}

}
