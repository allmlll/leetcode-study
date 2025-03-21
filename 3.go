package main

func lengthOfLongestSubstring(s string) int {
	res := 0
	j := -1
	n := len(s)
	mp := make(map[byte]int)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(mp, s[i-1])
		}
		for j+1 < n && mp[s[j+1]] == 0 {
			mp[s[j+1]]++
			j++
		}
		res = max(res, j-i+1)
	}
	return res
}
