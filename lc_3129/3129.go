package lc_3129

func numberOfStableArrays(zero int, one int, limit int) int {
	var results [][]int
	var res = 0

	generatePermutations(zero, one, []int{}, &results)
	if limit <= 1 {
	loop:
		for _, result := range results {

			for i := 0; i < zero+one-1; i++ {
				var nums [2]int
				for j := i; j <= i+1 && i+1 <= len(result)-1; j++ {
					if result[j] == 0 {
						nums[0]++
					} else {
						nums[1]++
					}

				}
				if nums[0] == 0 || nums[1] == 0 {
					continue loop

				}

			}
			res++

		}
	}
	if limit > zero+one {
		res = len(results)

	}
	if limit > 1 {
	loop2:
		for _, result := range results {
			for i := 0; i < zero+one; i++ {
				var nums [2]int
				for j := i; j <= i+limit && i+limit <= len(result)-1; j++ {
					if result[j] == 0 {
						nums[0]++
					} else {
						nums[1]++
					}

				}
				if nums[0] == 0 || nums[1] == 0 {
					continue loop2
				}
			}
			res++

		}
	}

	return res
}
func generatePermutations(n int, m int, prefix []int, results *[][]int) {
	// 基本情况：如果 n 和 m 都为 0，则已经生成了一个完整的排列
	if n == 0 && m == 0 {
		// 深拷贝 prefix 到一个新的切片中，避免引用问题
		result := make([]int, len(prefix))
		copy(result, prefix)
		*results = append(*results, result)
		return
	}

	// 如果还可以添加 0，那么就添加一个 0 并递归
	if n > 0 {
		generatePermutations(n-1, m, append(prefix, 0), results)
	}

	// 如果还可以添加 1，那么就添加一个 1 并递归
	if m > 0 {
		generatePermutations(n, m-1, append(prefix, 1), results)
	}
}

//func main() {
//	fmt.Println(numberOfStableArrays(2, 3, 2))
//}
