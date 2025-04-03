package lc_560

func subarraySum(nums []int, k int) (count int) {
	mp := make(map[int]int)
	//1,2,3
	sum := 0
	mp[0] = 1
	for _, num := range nums {
		sum += num

		if _, ok := mp[sum-k]; ok {
			count += mp[sum-k]
		}
		mp[sum]++
	}
	return count
}
