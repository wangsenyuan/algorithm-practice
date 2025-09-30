package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, ans, a := drive(reader)
	fmt.Println(ans)
	if ans == "Possible" {
		s := fmt.Sprintf("%v", a)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) (f []int, b []int, ans string, a []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	f = make([]int, n)
	b = make([]int, m)
	for i := range n {
		fmt.Fscan(reader, &f[i])
	}
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	ans, a = solve(m, f, b)
	return
}

func solve(m int, f []int, b []int) (string, []int) {
	n := len(f)
	// n <= m
	freq := make([]int, n+1)
	pos := make([]int, n+1)
	for i, v := range f {
		freq[v]++
		pos[v] = i + 1
	}
	// b[i] = f[a[i]] => a[i] = ?
	a := make([]int, m)

	ambiguous := false

	for i := 0; i < m; i++ {
		x := b[i]
		if freq[x] == 0 {
			return "Impossible", nil
		}
		if freq[x] > 1 {
			ambiguous = true
			continue
		}
		a[i] = pos[x]
	}
	if ambiguous {
		return "Ambiguity", nil
	}

	return "Possible", a
}
