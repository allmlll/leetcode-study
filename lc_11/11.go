package lc_11

func maxArea(height []int) (res int) {
	left := 0
	right := len(height) - 1
	res = 0
	for {

		if left > right {
			return
		}

		h := max(height[left], height[right])
		cres := h * (right - left)
		if cres > res {
			res = cres
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}

	}
}
