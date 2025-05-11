package p3541

func maxFreqSum(s string) int {
	freq := make([]int, 26)
	for _, x := range []byte(s) {
		freq[x-'a']++
	}
	ans := []int{0, 0}
	for i := 0; i < 26; i++ {
		if isVowel(byte(i + 'a')) {
			ans[0] = max(ans[0], freq[i])
		} else {
			ans[1] = max(ans[1], freq[i])
		}
	}
	return ans[0] + ans[1]
}

func isVowel(x byte) bool {
	return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u'
}
