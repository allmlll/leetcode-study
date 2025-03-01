package main

func closetTarget(words []string, target string, startIndex int) int {
	return min(left(words, target, startIndex), right(words, target, startIndex))
}
func left(words []string, target string, startIndex int) int {
	res := 0
	for i := 0; i < len(words); i++ {

		if words[startIndex] == target {
			return res
		}
		startIndex--
		res++
		if startIndex < 0 {
			startIndex = len(words) - 1
		}
	}
	return -1
}
func right(words []string, target string, startIndex int) int {
	res := 0
	for i := 0; i < len(words); i++ {

		if words[startIndex] == target {
			return res
		}
		startIndex++
		res++
		if startIndex > len(words)-1 {
			startIndex = 0
		}
	}
	return -1
}
