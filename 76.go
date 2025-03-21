package main

import "math"

func minWindow(s string, t string) string {
	//两个map一个记录t，另一个记录当前遍历数量
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	//result
	len := math.MaxInt32
	ansL, ansR := -1, -1
	//检查是否包含
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	//
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func minWindow2(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	slen := len(s)
	//初始化结果集
	len := math.MaxInt32
	lres, rres := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < slen; r++ {

		//去除无效字符,r右移
		if ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		//check    abbbbbbbcaabc
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				lres, rres = l, r

			}
			cnt[s[l]]--
			l++

		}

	}

	if lres == -1 {
		return ""
	}

	return s[lres : rres+1]

}
func minWindow3(s string, p string) string {
	if len(s) < len(p) {
		return ""
	}

	min_len := math.MaxInt

	cnt_p := make([]int, 128)
	for _, ch := range p {
		cnt_p[ch]++
	}

	vaild := 0
	cnt_s := make([]int, 128)
	l, start := 0, 0
	for r, ch := range s {
		cnt_s[ch]++
		if cnt_s[ch] <= cnt_p[ch] {
			vaild++
		}

		for vaild == len(p) {
			if cur_len := r - l + 1; cur_len < min_len {
				min_len = cur_len
				start = l
			}

			chL := s[l]

			cnt_s[chL]--
			if cnt_s[chL] < cnt_p[chL] {
				vaild--
			}
			l++
		}

	}

	if min_len == math.MaxInt {
		return ""
	}

	return s[start : start+min_len]
}
