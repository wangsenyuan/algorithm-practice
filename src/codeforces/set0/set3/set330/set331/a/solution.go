package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	best, cut := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %d\n", best, len(cut)))
	for _, x := range cut {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) (best int, cut []int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	best, cut = solve(a)
	return
}

func solve(a []int) (best int, cut []int) {
	first := make(map[int]int)

	n := len(a)
	sum := make([]int, n+1)

	for i, v := range a {
		sum[i+1] = sum[i] + max(v, 0)
		if _, ok := first[v]; !ok {
			first[v] = i
		}
	}
	best = -inf
	var pick []int

	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if j, ok := first[v]; ok && j < i {
			tmp := sum[i] - sum[j+1] + 2*v
			if tmp > best {
				best = tmp
				pick = []int{j, i}
			}
		}
	}

	for i := 0; i < pick[0]; i++ {
		cut = append(cut, i+1)
	}

	for i := pick[0] + 1; i < pick[1]; i++ {
		if a[i] < 0 {
			cut = append(cut, i+1)
		}
	}

	for i := pick[1] + 1; i < n; i++ {
		cut = append(cut, i+1)
	}

	return
}

const inf = 1 << 60
