package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var s, m int
		fmt.Fscan(reader, &s, &m)
		res := solve(s, m)
		fmt.Fprintln(writer, res)
	}
}

func solve(s int, m int) int {
	h := bits.Len(uint(m))

	check := func(n int) bool {
		t := s
		for d := h - 1; d >= 0; d-- {
			if (m>>d)&1 == 1 {
				w := min(n, t/(1<<d))
				t -= w * (1 << d)
			}
		}

		return t == 0
	}
	inf := s + 1
	ans := sort.Search(inf, check)
	if ans == inf {
		return -1
	}
	return ans
}
