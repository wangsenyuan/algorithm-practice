package p3921

import "strconv"

func scoreValidator(events []string) []int {
	var score int
	var count int
	for _, cur := range events {
		if cur == "W" {
			count++
		} else if cur == "WD" || cur == "NB" {
			score++
		} else {
			v, _ := strconv.Atoi(cur)
			score += v
		}
		if count == 10 {
			break
		}
	}
	return []int{score, count}
}
