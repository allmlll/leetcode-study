package main

func longestConsecutive(nums []int) (res int) {
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	res = 0
	for num := range mp {
		if !mp[num-1] {
			cnum := num
			k := 1
			for mp[cnum+1] {
				cnum++
				k++
			}
			if res < k {
				res = k
			}
		}

	}
	return res
}
