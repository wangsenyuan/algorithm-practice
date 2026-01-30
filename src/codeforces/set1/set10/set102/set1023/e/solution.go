package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readString(reader)
	n, _ := strconv.Atoi(first)

	ask := func(res []int) bool {
		fmt.Printf("? %d %d %d %d\n", res[0], res[1], res[2], res[3])
		s := readString(reader)
		return s == "YES"
	}

	ans := solve(n, ask)
	fmt.Printf("! %s\n", ans)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(n int, ask func([]int) bool) string {

	var buf []byte
	var r, c int
	for (n-1-r)+(n-1-c) > n-1 {
		// prefer down
		if r+1 < n {
			ok := ask([]int{r + 2, c + 1, n, n})
			if ok {
				buf = append(buf, 'D')
				r++
				continue
			}
		}
		// go right
		buf = append(buf, 'R')
		c++
	}
	// try reach r, c from backend
	var suf []byte
	r1, c1 := n-1, n-1

	// r, c is the most bottom-left one
	// when r1 == r, only R is needed
	for r1 > r {
		// c1 - 1 + 1
		// prefer go left
		ok := ask([]int{1, 1, r1 + 1, c1})
		if ok {
			suf = append(suf, 'R')
			c1--
		} else {
			// have to go up
			suf = append(suf, 'D')
			r1--
		}
	}

	// r == r1 holds
	for c < c1 {
		buf = append(buf, 'R')
		c++
	}

	slices.Reverse(suf)
	buf = append(buf, suf...)
	return string(buf)
}
