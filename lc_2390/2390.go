package lc_2390

func removeStars(s string) string {

	var lens = len(s)
	index := make([]int, lens)
	count := 0
	k := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			index[i] = 1
			if count == 0 {
				for j := i - 1; j >= 0; j-- {
					if index[j] == 0 {
						count++
						index[j] = 1
						k = j
						break
					}
					if j == 0 {
						index[i] = 0
					}
				}
			} else {
				for j := k - 1; j >= 0; j-- {
					if index[j] == 0 {
						count++
						index[j] = 1
						k = j
						break
					}
					if j == 0 {
						index[i] = 0
					}
				}

			}

		}

	}
	var ss string
	for k := 0; k < len(s); k++ {
		if index[k] == 0 {
			ss = ss + string(s[k])
		}
	}
	return ss

}

//func main() {
//	fmt.Println(removeStars("leet**cod*e"))
//}
