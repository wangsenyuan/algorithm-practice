package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprint(writer, drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	var sum int
	var sum1 int
	var hi, lo int
	for i, v := range a {
		sum += v
		if i&1 == 0 {
			sum1 += v
		} else {
			sum1 -= v
		}
		hi = max(hi, sum1)
		lo = min(lo, sum1)
	}

	m := hi - lo + 1
	lines := make([][]byte, m)
	for i := range lines {
		lines[i] = make([]byte, sum)
		for j := range sum {
			lines[i][j] = ' '
		}
	}

	play := func(x0 int, y0 int, dx int, dy int) {
		// abs(dx) = abs(dy)
		var ch byte = '/'
		w := 1
		j := y0 - lo
		if dy < 0 {
			ch = '\\'
			w = -1
			j--
		}
		for i := x0; i < x0+dx; i++ {
			lines[m-j-1][i] = ch
			j += w
		}
	}

	var x, y int
	for i, v := range a {
		if i&1 == 0 {
			play(x, y, v, v)
			x += v
			y += v
		} else {
			play(x, y, v, -v)
			x += v
			y -= v
		}
	}

	var ans bytes.Buffer
	for _, cur := range lines[1:] {
		ans.WriteString(string(cur))
		ans.WriteByte('\n')
	}
	return ans.String()
}
