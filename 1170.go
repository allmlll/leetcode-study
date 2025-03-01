package main

import "sort"

func findSmallestChar(s string) rune {

	smallest := rune(s[0]) // 假设第一个字符是最小的

	for _, char := range s[1:] {
		if char < smallest {
			smallest = char // 找到更小的字符则更新smallest
		}
	}

	return smallest
}
func countCharOccurrences(s string, c rune) int {
	count := 0
	for _, char := range s {
		if char == c {
			count++
		}
	}
	return count
}
func numSmallerByFrequency(queries []string, words []string) []int {

	qus := make([]int, len(queries))
	end := make([]int, len(queries))
	wos := make([]int, len(words))
	for i, query := range queries {
		ch := findSmallestChar(query)
		qus[i] = countCharOccurrences(query, ch)
	}
	for i, word := range words {
		ch := findSmallestChar(word)
		wos[i] = countCharOccurrences(word, ch)
	}
	sort.Ints(wos)
	an := 0
	for i := 0; i < len(qus); i++ {
		for j := 0; j < len(wos); j++ {
			if qus[i] < wos[j] {
				an = len(wos) - j
				end[i] = an
				break
			}
		}
	}
	return end

}
