package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	ans := solve(n)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(ans)))
	for _, cur := range ans {
		for i, x := range cur {
			buf.WriteString("{")
			for j, y := range x {
				buf.WriteString(fmt.Sprintf("%d", y))
				if j+1 < len(x) {
					buf.WriteString(",")
				}
			}
			buf.WriteString("}")
			if i+1 < len(cur) {
				buf.WriteString(",")
			} else {
				buf.WriteString("\n")
			}
		}
	}
	buf.WriteTo(os.Stdout)
}

func solve(n int) [][][]int {

	v := [][]int{{0}}

	for i := 2; i <= n; i++ {
		var u [][]int
		for j := range len(v) {
			cur := slices.Clone(v[j])
			var m int
			for _, k := range cur {
				m = max(m, k+1)
			}
			if j%2 == 0 {
				cur = append(cur, 0)
				u = append(u, slices.Clone(cur))
				cur = cur[:len(cur)-1]
				for x := 2; x <= m; x++ {
					cur = append(cur, x)
					u = append(u, slices.Clone(cur))
					cur = cur[:len(cur)-1]
				}
				cur = append(cur, 1)
				u = append(u, slices.Clone(cur))
				cur = cur[:len(cur)-1]
			} else {
				cur = append(cur, 1)
				u = append(u, slices.Clone(cur))
				cur = cur[:len(cur)-1]
				for x := 2; x <= m; x++ {
					cur = append(cur, x)
					u = append(u, slices.Clone(cur))
					cur = cur[:len(cur)-1]
				}
				cur = append(cur, 0)
				u = append(u, slices.Clone(cur))
				cur = cur[:len(cur)-1]
			}
		}
		v = u
	}
	var res [][][]int

	for _, cur := range v {
		var mx int
		for _, j := range cur {
			mx = max(mx, j)
		}
		var tmp [][]int
		for j := 0; j <= mx; j++ {
			var buf []int
			for k := 0; k < len(cur); k++ {
				if cur[k] == j {
					buf = append(buf, k+1)
				}
			}
			tmp = append(tmp, buf)
		}
		res = append(res, tmp)
	}

	return res
}

func copy2d(a [][]int) [][]int {
	b := make([][]int, len(a))
	for i := range a {
		b[i] = make([]int, len(a[i]))
		copy(b[i], a[i])
	}
	return b
}
