package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(reader, &n, &s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) int {
	n := len(s)
	freq := make(BIT, 4*n)
	offset := 2 * n
	freq.add(offset, 1)
	var cnt int
	var res int
	for _, x := range s {
		switch x {
		case 'A':
			cnt++
		case 'B':
			cnt--
		}
		// cnt
		res += freq.get(offset + cnt - 1)
		freq.add(offset+cnt, 1)
	}

	return res
}

type BIT []int

func (b BIT) add(p int, v int) {
	p++
	for p < len(b) {
		b[p] += v
		p += p & -p
	}
}

func (b BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += b[p]
		p -= p & -p
	}
	return res
}
