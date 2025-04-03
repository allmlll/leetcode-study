package lc_2134

func minSwaps(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sum := 0
	first := 0
	flag := 0
	lenth := 0
	for i := 0; i < 2*len(nums); i++ {
		j := i
		j = j % len(nums)
		if nums[j] == 1 && j > 0 && nums[(j-1)%len(nums)] == 0 {
			first = i
			flag = 1
		}
		if nums[j] == 0 && flag == 1 {
			lenth = max(lenth, i-first)
			flag = 0
		}
		if nums[j] == 1 && i < len(nums) {
			sum++

		}
	}

	return sum - lenth

}

//func main() {
//	fmt.Println(minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
//}
