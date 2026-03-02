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
	var w, m, k int
	fmt.Fscan(reader, &w, &m, &k)
	return solve(w, m, k)
}

func solve(w int, m int, k int) int {
	// s(n) = digits count of n
	// sum(s(n)) * k <= w

	count := func(r int) int {
		if r == 0 {
			return 0
		}
		var res int
		x := 1
		d := 1
		for x <= r {
			y := min(r, x*10-1)
			res += (y - x + 1) * d
			x = y + 1
			d++
		}
		return res
	}

	base := count(m - 1)

	ds := count(m) - base
	if ds > w/k || ds*k > w {
		return 0
	}

	check := func(n int) bool {
		r := m + n
		// w...r的位数
		tmp := count(r) - base
		if tmp > w/k || tmp*k > w {
			return true
		}
		return false
	}

	return sort.Search(1e16, check)
}
