package lc_49

import "sort"

func groupAnagrams(strs []string) (res [][]string) {
	mp := make(map[string][]string)
	for _, str := range strs {
		r := []rune(str)
		sort.Slice(r, func(i, j int) bool {
			if r[i] > r[j] {
				return true
			}
			return false
		})
		mp[string(r)] = append(mp[string(r)], str)
	}
	for _, m := range mp {
		res = append(res, m)
	}
	return res
}
