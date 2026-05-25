package p3941

import "slices"

func passwordStrength(password string) int {
	buf := []byte(password)
	slices.Sort(buf)
	var score int
	for i := 0; i < len(buf); {
		j := i
		for i < len(buf) && buf[i] == buf[j] {
			i++
		}
		if buf[j] >= 'a' && buf[j] <= 'z' {
			score++
		} else if buf[j] >= 'A' && buf[j] <= 'Z' {
			score += 2
		} else if buf[j] >= '0' && buf[j] <= '9' {
			score += 3
		} else {
			score += 5
		}
	}
	return score
}
