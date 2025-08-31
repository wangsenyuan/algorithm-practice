package p3663

func getLeastFrequentDigit(n int) int {
	freq := make([]int, 10)
	for n > 0 {
		freq[n%10]++
		n /= 10
	}
	res := 10
	for i := 0; i < 10; i++ {
		if freq[i] > 0 && (res == 10 || freq[i] < freq[res]) {
			res = i
		}
	}
	return res
}
