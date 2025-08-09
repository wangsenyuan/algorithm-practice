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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)

	pos := make([]int, n+1)
	for i, v := range p {
		pos[v] = i
	}

	cnt := make(BIT, n+1)

	var res int

	for i := 1; i <= n; i++ {
		j := pos[i]
		l := j - cnt.pre(j+1)
		r := n - i - l
		res += min(l, r)
		cnt.update(j+1, 1)
	}
	return res
}

type BIT []int

func (f BIT) update(i int, v int) {
	for i < len(f) {
		f[i] += v
		i += i & -i
	}
}

func (f BIT) pre(i int) (res int) {
	for i > 0 {
		res += f[i]
		i -= i & -i
	}
	return res
}
