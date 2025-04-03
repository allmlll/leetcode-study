package lc_1334

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	res := n
	count := 0

	for i := 0; i < n; i++ {
		t := distanceThreshold
		dfs(edges, i, t, count)
		res = min(count, res)
	}
	return res
}
func dfs(edges [][]int, n int, distanceThreshold int, count int) {
	for _, edge := range edges {
		t := distanceThreshold
		if edge[0] == n {
			t -= edge[2]
			if t <= 0 {
				continue
			}
			count++
			dfs(edges, edges[n][1], t, count)
		}
	}
}

//func main() {
//	edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
//	fmt.Println(findTheCity(4, edges, 4))
//}
