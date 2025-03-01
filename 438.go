package main

func findAnagrams(s string, p string) []int {
	var res []int
	mp := map[rune]int{}
	mp2 := map[rune]int{}
	for _, ps := range p {
		mp[ps]++
	}

loop:
	for i := 0; i <= len(s)-len(p); i++ {
		for r, i2 := range mp {
			mp2[r] = i2
		}
		for j := i; j < i+len(p); j++ {
			for r, _ := range mp2 {
				if rune(s[j]) == r {
					mp2[r]--
					break
				}
			}
		}

		for _, mpst := range mp2 {
			if mpst != 0 {
				continue loop
			}
		}
		res = append(res, i)

	}
	return res
}
