package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, a, d int
	fmt.Fscan(reader, &n, &m, &a, &d)
	t := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &t[i])
	}
	return solve(n, a, d, t)
}

func solve(n int, a int, d int, t []int) int {
	// deduplicate consecutive equal times (simultaneous arrivals = one event)
	t = slices.Compact(t)

	res := 0
	emp := 1        // next employee index (1-based, arrives at emp*a)
	closeTime := -1 // time when door will close; -1 means door is closed

	// processEmps handles all employee arrivals up to time `deadline`
	processEmps := func(deadline int) {
		if emp > n {
			return
		}
		lastEmp := min(n, deadline/a)
		if lastEmp < emp {
			return
		}

		// skip employees that arrive while door is still open
		if emp*a <= closeTime {
			firstNew := closeTime/a + 1
			if firstNew > lastEmp {
				// all employees in range enter through already-open door
				emp = lastEmp + 1
				return
			}
			emp = firstNew
		}

		// door is closed for employees emp..lastEmp
		// each opening lets batch = floor(d/a) + 1 employees through
		count := lastEmp - emp + 1
		batch := d/a + 1
		openings := (count + batch - 1) / batch
		res += openings

		// the last opening starts at this employee
		lastOpenEmp := emp + (openings-1)*batch
		closeTime = lastOpenEmp*a + d
		emp = lastEmp + 1
	}

	for i := range len(t) {
		processEmps(t[i])
		if t[i] > closeTime {
			res++
			closeTime = t[i] + d
		}
	}

	// process remaining employees after last client
	processEmps(n * a)

	return res
}
