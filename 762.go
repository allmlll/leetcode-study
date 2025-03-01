package main

func countPrimeSetBits(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		if isPrime(countSetBits(i)) {
			res++
		}

	}
	return res
}
func countSetBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1 // Clear the least significant bit set
		count++
	}
	return count
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Only check for factors up to the square root of n
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
