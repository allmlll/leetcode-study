package lc_54

func spiralOrder(matrix [][]int) []int {
	if len(matrix[0]) == 0 || len(matrix) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	order := make([]int, columns*rows)
	index := 0
	left, right, top, bottom := 0, columns-1, 0, rows-1

	for left <= right && top <= bottom {

		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][left]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column >= left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom + 1; row >= top; row-- {
				order[index] = matrix[row][right]
				index++
			}
		}

		left++
		right--
		top++
		bottom--

	}
	return order

}
