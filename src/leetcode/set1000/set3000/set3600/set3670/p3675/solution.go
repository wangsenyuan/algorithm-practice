package p3675

func minOperations(s string) int {
	freq := make([]int, 27)
	for i := range s {
		x := int(s[i] - 'a')
		freq[x]++
	}
	var res int
	for i := 1; i < 26; i++ {
		if freq[i] > 0 {
			res++
		}
		freq[i+1] += freq[i]
	}
	return res
}
