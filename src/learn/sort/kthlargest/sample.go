package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 9, 8, 7}
	fmt.Println(secondLargest(nums))
}

func secondLargest(nums []int) int {
	n := len(nums)
	if n < 2 {
		panic("insufficient number to find the second largest number")
	}

	candidates := make([]*Candidate, n)
	for i := range candidates {
		candidates[i] = &Candidate{nums[i], make([]int, 0)}
	}

	for len(candidates) > 1 {
		i := 0
		j := 0
		for i+1 < len(candidates) {
			a, b := candidates[i], candidates[i+1]
			if a.val >= b.val {
				a.less = append(a.less, b.val)
				candidates[j] = a
			} else {
				b.less = append(b.less, a.val)
				candidates[j] = b
			}
			i += 2
			j++
		}

		if i < len(candidates) {
			// Compare the last candidate with the current winner
			if j > 0 {
				last := candidates[i]
				winner := candidates[j-1]
				if last.val >= winner.val {
					last.less = append(last.less, winner.val)
					candidates[j-1] = last
				} else {
					winner.less = append(winner.less, last.val)
				}
			} else {
				candidates[j] = candidates[i]
				j++
			}
		}
		candidates = candidates[:j]
	}

	largest := candidates[0]
	less := largest.less

	second := less[0]

	for i := 1; i < len(less); i++ {
		if less[i] > second {
			second = less[i]
		}
	}
	return second
}

type Candidate struct {
	val  int
	less []int
}
