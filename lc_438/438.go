package lc_438

func findAnagrams(s string, p string) []int {
	var res []int
	mp := make(map[byte]int)
	plen := len(p)
	slen := len(s)
	if plen > slen {
		return []int{}
	}
	block := make(map[byte]int)
	for i := 0; i < plen; i++ {
		mp[p[i]]++
	}
	//初始化滑块
	for i := 0; i < plen; i++ {
		block[s[i]]++
	}
	if check(block, mp) {
		res = append(res, 0)
	}
	for i := plen; i < slen; i++ {

		if _, ok := mp[s[i]]; ok {
			block[s[i-plen]]--
			block[s[i]]++
		} else {
			block = make(map[byte]int)
			if i+plen < slen {
				for j := i + 1; j < i+1+plen; j++ {
					block[s[j]]++

				}

			}
			i = i + plen

		}
		if check(block, mp) {
			res = append(res, i-plen+1)
		}
	}
	return res
}
func check(block, mp map[byte]int) bool {
	for key, value := range mp {
		if block[key] != value {
			return false
		}
	}
	return true
}
