package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, s int
	fmt.Fscan(reader, &n, &s)
	best, p := solve(n, s)
	fmt.Println(best)
	if best < 0 {
		return
	}
	for i := range n {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for _, v := range p {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func solve(n int, s int) (best int, p []int) {
	sum := (n + 1) * n / 2
	if s < sum {
		return -1, nil
	}
	p = make([]int, n)
	for i := range n {
		p[i] = i + 1
	}

	d := s - sum

	for i, j := 0, n-1; i < j && d > 0; i++ {
		x := min(d, j-i)
		j = i + x
		d -= x
		p[i], p[j] = p[j], p[i]
		j--
	}

	best = s - d
	return
}
