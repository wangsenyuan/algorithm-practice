package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, m, g, h := drive(reader)
	fmt.Println(m)
	if m > 0 {
		s := fmt.Sprintf("%v", g)
		fmt.Println(s[1 : len(s)-1])
		s = fmt.Sprintf("%v", h)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) (f []int, m int, g []int, h []int) {
	var n int
	fmt.Fscan(reader, &n)
	f = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &f[i])
	}
	m, g, h =  solve(slices.Clone(f))
	return
}

func solve(f []int) (m int, g []int, h []int) {
	n := len(f)

	for i := range n {
		f[i]--
	}

	g = make([]int, n)

	for i := range n {
		if f[f[i]] != f[i] {
			return -1, nil, nil
		}
		if f[i] == i {
			m++
			g[i] = len(h)
			h = append(h, i)
		}
	}

	for i := range n {
		g[i] = g[f[i]]
	}

	for i := range n {
		g[i]++
	}

	for i := range m {
		h[i]++
	}


	return
}