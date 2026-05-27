package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, a, b, c int
	fmt.Fscan(reader, &n, &a, &b, &c)
	t := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &t[i])
	}
	return solve(a, b, c, t)
}

func solve(a int, b int, c int, t []int) int {
	sort.Ints(t)
	a, b, c = sortThree(a, b, c)

	n := len(t)
	if countGreater(t, a+b+c) > 0 {
		return -1
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if canFinish(t, a, b, c, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func canFinish(t []int, a int, b int, c int, hours int) bool {
	n := len(t)
	ab := a + b
	ac := a + c
	bc := b + c

	gt := func(x int) int {
		return countGreater(t, x)
	}

	// These strongest criminals force the corresponding full-hour patterns.
	needAll := gt(bc)
	if needAll > hours {
		return false
	}

	needBC := gt(ac) - gt(bc)         // pair b+c, leaves a free
	needAC := gt(max(ab, c)) - gt(ac) // pair a+c, leaves b free
	needAB := gt(c) - gt(max(ab, c))  // pair a+b, leaves c free
	needPair := needBC + needAC + needAB
	remainingHours := hours - needAll
	if needPair > remainingHours {
		return false
	}

	freeHours := remainingHours - needPair
	remainingCriminals := n - gt(c)

	// Among the free hours, let h be the number of hours used as (a+b) + c.
	// The other freeHours-h hours are used as a + b + c.
	lo, hi := 0, freeHours

	checkThreshold := func(th int, need int) bool {
		beta, coef := 0, 0
		if a > th {
			beta += needBC + freeHours
			coef--
		}
		if b > th {
			beta += needAC + freeHours
			coef--
		}
		if c > th {
			beta += needAB + freeHours
		}
		if ab > th {
			coef++
		}

		if coef == 0 {
			return beta >= need
		}
		if coef > 0 {
			lo = max(lo, divCeil(need-beta, coef))
			return true
		}
		hi = min(hi, divFloor(beta-need, -coef))
		return true
	}

	thresholds := []int{-1, a, b, c, ab}
	for _, th := range thresholds {
		need := remainingCriminals
		if th >= 0 {
			if th >= c {
				need = 0
			} else {
				need = gt(th) - gt(c)
			}
		}
		if !checkThreshold(th, need) {
			return false
		}
	}

	return lo <= hi && hi >= 0 && lo <= freeHours
}

func countGreater(arr []int, x int) int {
	return len(arr) - sort.SearchInts(arr, x+1)
}

func divCeil(a int, b int) int {
	if a <= 0 {
		return 0
	}
	return (a + b - 1) / b
}

func divFloor(a int, b int) int {
	if a < 0 {
		return -1
	}
	return a / b
}

func sortThree(a int, b int, c int) (x int, y int, z int) {
	x = min(a, b, c)
	z = max(a, b, c)
	y = a + b + c - x - z
	return
}
