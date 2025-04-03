package lc_240

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])

	x, y := 0, m-1
	for x <= n-1 && y >= 0 {

		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}

	}
	return false
}
