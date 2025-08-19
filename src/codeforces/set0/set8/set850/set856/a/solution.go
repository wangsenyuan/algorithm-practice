package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		ans := drive(reader)
		if len(ans) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, v := range ans {
				buf.WriteString(fmt.Sprintf("%d ", v))
			}
			buf.WriteString("\n")
		}
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

var bad []bool

func init() {
	bad = make([]bool, 1000001)
}

func solve(a []int) []int {
	// a[i] + b[?] 得到的 n * n 个数，要各不相同
	sort.Ints(a)

	clear(bad)

	n := len(a)

	var b []int
	for x := 1; x <= 1e6 && len(b) < n; x++ {
		if bad[x] {
			continue
		}
		b = append(b, x)
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				bad[x-a[j]+a[i]] = true
			}
		}
	}
	return b
}
