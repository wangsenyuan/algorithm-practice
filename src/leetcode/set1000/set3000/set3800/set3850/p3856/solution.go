package p3856

func trimTrailingVowels(s string) string {
	for len(s) > 0 && isVowel(s[len(s)-1]) {
		s = s[:len(s)-1]
	}
	return s
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
