package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for range tc {
		res := process(reader)
		buf.WriteString(output(res))
	}
	buf.WriteTo(os.Stdout)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func process(reader *bufio.Reader) [][]int {
	n := readNum(reader)
	q := readNum(reader)
	qs := make([][]int, q)
	for i := range q {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '-' {
			qs[i] = make([]int, 2)
		} else {
			qs[i] = make([]int, 1)
		}
		pos := 3
		for j := range qs[i] {
			pos = readInt(s, pos, &qs[i][j]) + 1
		}
	}

	return solve(n, qs)
}

func output(ans [][]int) string {
	var buf bytes.Buffer
	for _, cur := range ans {
		for _, x := range cur {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.Truncate(buf.Len() - 1)
		buf.WriteByte('\n')
	}
	return buf.String()
}

func solve(n int, queries [][]int) [][]int {

	ans := make([][]int, len(queries))

	for i, cur := range queries {
		if len(cur) == 2 {
			ans[i] = []int{query1(n, cur[0]-1, cur[1]-1)}
		} else {
			ans[i] = query2(n, cur[0]-1)
		}
	}
	return ans
}

func query1(n int, x int, y int) int {
	var loop func(x0 int, y0 int, x1 int, y1 int, lo int, hi int) int

	loop = func(x0 int, y0 int, x1 int, y1 int, lo int, hi int) int {
		if x0+1 == x1 && y0+1 == y1 {
			return lo + 1
		}
		mx := (x0 + x1) / 2
		my := (y0 + y1) / 2
		cnt := (hi - lo) / 4
		if x < mx && y < my {
			return loop(x0, y0, mx, my, lo, lo+cnt)
		}
		if x >= mx && y >= my {
			return loop(mx, my, x1, y1, lo+cnt, lo+2*cnt)
		}
		if x >= mx && y < my {
			return loop(mx, y0, x1, my, lo+2*cnt, lo+3*cnt)
		}
		return loop(x0, my, mx, y1, lo+3*cnt, hi)
	}

	return loop(0, 0, 1<<n, 1<<n, 0, 1<<(2*n))
}

func query2(n int, x int) []int {
	var loop func(x0 int, y0 int, x1 int, y1 int, lo int, hi int) []int

	loop = func(x0 int, y0 int, x1 int, y1 int, lo int, hi int) []int {
		if x0+1 == x1 && y0+1 == y1 {
			return []int{x0 + 1, y0 + 1}
		}
		mx := (x0 + x1) / 2
		my := (y0 + y1) / 2
		cnt := (hi - lo) / 4
		if x < lo+cnt {
			return loop(x0, y0, mx, my, lo, lo+cnt)
		}
		if x < lo+2*cnt {
			return loop(mx, my, x1, y1, lo+cnt, lo+2*cnt)
		}
		if x < lo+3*cnt {
			return loop(mx, y0, x1, my, lo+2*cnt, lo+3*cnt)
		}
		return loop(x0, my, mx, y1, lo+3*cnt, hi)
	}

	return loop(0, 0, 1<<n, 1<<n, 0, 1<<(2*n))
}
