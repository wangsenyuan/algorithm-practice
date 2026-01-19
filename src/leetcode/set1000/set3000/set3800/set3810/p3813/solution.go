package p3813

func vowelConsonantScore(s string) int {
	var vowels = "aeiou"

	isVowel := func(c rune) bool {
		for _, v := range vowels {
			if c == v {
				return true
			}
		}
		return false
	}

	cnt := make([]int, 2)
	for _, c := range s {
		if isVowel(c) {
			cnt[0]++
		} else if c >= 'a' && c <= 'z' {
			cnt[1]++
		}
	}

	if cnt[1] == 0 {
		return 0
	}
	return cnt[0] / cnt[1]
}
