package main

import (
	"bufio"
	"cmp"
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
		ans := drive(reader)
		s := fmt.Sprintf("%v", ans)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		a[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a [][]int) []int {
	n := len(a)
	var m int
	rows := make([]int, n)
	for i, v := range a {
		m = max(m, len(v))
		rows[i] = i
	}

	find := func(rows []int, d int) []int {
		for len(rows) > 1 {
			nextVal := 1 << 60
			var newRows []int
			for _, i := range rows {
				if len(a[i]) == d {
					return []int{i, d}
				}
				nextVal = min(nextVal, a[i][d])
			}
			for _, i := range rows {
				if a[i][d] == nextVal {
					newRows = append(newRows, i)
				}
			}
			rows = newRows
			d++
		}
		return []int{rows[0], len(a[rows[0]])}
	}

	slices.SortFunc(rows, func(i int, j int) int {
		return cmp.Or(len(a[i])-len(a[j]), i-j)
	})

	ans := make([]int, m)

	var pos int
	for d := 0; d < m; {
		for pos < n && len(a[rows[pos]]) <= d {
			pos++
		}
		// 找到这一列的最小值所在的行，可能有多少个
		tmp := find(rows[pos:], d)
		r, d1 := tmp[0], tmp[1]

		for i := d; i < d1; i++ {
			ans[i] = a[r][i]
		}

		d = d1
	}

	return ans
}
