package p3848

func isDigitorialPermutation(n int) bool {
	f := make([]int, 10)
	f[0] = 1
	for i := 1; i < 10; i++ {
		f[i] = f[i-1] * i
	}

	var sum int
	for num := n; num > 0; num /= 10 {
		x := num % 10
		sum += f[x]
	}

	freq := make([]int, 10)
	for n > 0 && sum > 0 {
		freq[n%10]++
		freq[sum%10]--
		n /= 10
		sum /= 10
	}

	if n != 0 || sum != 0 {
		// 长度不一样
		return false
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
