package p3598

func longestCommonPrefix(words []string) []int {
	n := len(words)
	if n == 1 {
		return []int{0}
	}
	pref := make([]int, n)
	for i := 1; i < n; i++ {
		a := words[i-1]
		b := words[i]
		for pref[i] < len(a) && pref[i] < len(b) && a[pref[i]] == b[pref[i]] {
			pref[i]++
		}
		pref[i] = max(pref[i], pref[i-1])
	}
	ans := make([]int, n)
	ans[n-1] = pref[n-2]
	var suf int
	for i := n - 2; i > 0; i-- {
		ans[i] = max(pref[i-1], suf)
		a := words[i]
		b := words[i+1]
		var j int
		for j < len(a) && j < len(b) && a[j] == b[j] {
			j++
		}
		suf = max(suf, j)
		j = 0
		a = words[i-1]
		for j < len(a) && j < len(b) && a[j] == b[j] {
			j++
		}
		ans[i] = max(ans[i], j)
	}
	ans[0] = suf
	return ans
}
