package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var a int
	var n int
	fmt.Fscan(reader, &a, &n)
	d := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &d[i])
	}
	return solve(a, d)
}

func solve(a int, d []int) int {
	if a == 0 {
		return d[0]
	}

	var ds []int
	for x := a; x > 0; x /= 10 {
		ds = append(ds, x%10)
	}
	slices.Reverse(ds)

	allowed := [10]bool{}
	for _, x := range d {
		allowed[x] = true
	}

	const inf = 1 << 60
	fill := func(x, digit, count int) int {
		for range count {
			if x > (inf-digit)/10 {
				return inf
			}
			x = x*10 + digit
		}
		return x
	}
	ans := inf
	relax := func(b int) {
		if b == inf {
			return
		}
		diff := a - b
		if diff < 0 {
			diff = -diff
		}
		ans = min(ans, diff)
	}

	minD, maxD := d[0], d[len(d)-1]
	if len(ds) > 1 {
		relax(fill(0, maxD, len(ds)-1))
	}
	for _, x := range d {
		if x > 0 {
			relax(fill(x, minD, len(ds)))
			break
		}
	}

	prefix := 0
	for i, x := range ds {
		lower, upper := -1, -1
		for _, y := range d {
			if y < x {
				lower = y
			} else if y > x && upper < 0 {
				upper = y
			}
		}

		rest := len(ds) - i - 1
		if lower >= 0 {
			relax(fill(prefix*10+lower, maxD, rest))
		}
		if upper >= 0 {
			relax(fill(prefix*10+upper, minD, rest))
		}
		if !allowed[x] {
			return ans
		}
		prefix = prefix*10 + x
	}
	return 0
}
