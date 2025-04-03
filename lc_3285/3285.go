package lc_3285

func stableMountains(height []int, threshold int) []int {
	var res []int
	for i := 0; i < len(height)-1; i++ {
		if height[i] > threshold {
			res = append(res, i+1)
		}
	}
	return res
}
