package lc_3142

func satisfiesConditions(grid [][]int) bool {
	if len(grid[0]) == 1 {
		return false
	}
	for i := 0; i < len(grid[0])-1; i++ {

		for j := 0; j < len(grid)-1; j++ {
			if grid[j][i] != grid[j+1][i] || grid[j][i] == grid[j][i+1] {
				return false
			}
		}
	}
	return true
}

//func main() {
//	//[[1,0,2],[1,0,2]]
//	gird := [][]int{{4, 4}, {4, 4}}
//	fmt.Println(satisfiesConditions(gird))
//}
