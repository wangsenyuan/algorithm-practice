package p3707

func scoreBalance(s string) bool {
	var sum int
	for i := range len(s) {
		sum += int(s[i] - 'a' + 1)
	}
	if sum%2 == 1 {
		return false
	}
	var pref int
	for i := range len(s) {
		x := int(s[i] - 'a' + 1)
		pref += x
		sum -= x
		if pref == sum {
			return true
		}
	}
	return false
}
