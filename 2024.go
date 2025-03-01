package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	res := 0
	j := 0
	ans := 0
	m := 0
	kk := k
	for i := 0; i < len(answerKey); i++ {

		if answerKey[i] != answerKey[j] {
			k--
			m = i
		}
		ans++
		if k == 0 {
			res = max(res, ans)
			ans = 0
			i = m
			k = kk
			j = m
		}
	}
	return res
}

//func main() {
//	fmt.Println(maxConsecutiveAnswers("TFFT", 1))
//}
