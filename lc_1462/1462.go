package lc_1462

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	var bools []bool
	var flag bool = false
	for i := 0; i < len(prerequisites); i++ {
		for j := 0; j < len(prerequisites); j++ {
			if prerequisites[i][1] == prerequisites[j][0] && prerequisites[i][0] != prerequisites[j][1] {
				prerequisites = append(prerequisites, []int{prerequisites[i][0], prerequisites[j][1]})
			}
		}
	}
	for _, query := range queries {
		flag = false
		for _, prerequisite := range prerequisites {
			if check(query, prerequisite) {
				flag = true
				break
			}

		}

		if !flag {
			bools = append(bools, false)
			continue
		}
		bools = append(bools, true)

	}
	return bools
}

func check(query, prerequisite []int) bool {
	if query[0] == prerequisite[0] && query[1] == prerequisite[1] {
		return true
	}
	return false
}

//func main() {
//	fmt.Println(checkIfPrerequisite(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}))
//}
