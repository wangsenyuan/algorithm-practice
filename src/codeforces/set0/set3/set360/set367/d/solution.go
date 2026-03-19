package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, d int
	fmt.Fscan(reader, &n, &m, &d)

	sets := make([][]int, m)
	for i := 0; i < m; i++ {
		var sz int
		fmt.Fscan(reader, &sz)
		sets[i] = make([]int, sz)
		for j := 0; j < sz; j++ {
			fmt.Fscan(reader, &sets[i][j])
		}
	}

	return solve(n, m, d, sets)
}

func solve(n int, m int, d int, sets [][]int) int {
	setOf := make([]int, n+1)
	for i, cur := range sets {
		for _, x := range cur {
			setOf[x] = i
		}
	}

	limit := 1 << m
	bad := make([]bool, limit)

	freq := make([]int, m)
	var mask int

	add := func(setID int) {
		if freq[setID] == 0 {
			mask |= 1 << setID
		}
		freq[setID]++
	}

	remove := func(setID int) {
		freq[setID]--
		if freq[setID] == 0 {
			mask ^= 1 << setID
		}
	}

	for i := 1; i <= d; i++ {
		add(setOf[i])
	}
	bad[mask] = true

	for r := d + 1; r <= n; r++ {
		remove(setOf[r-d])
		add(setOf[r])
		bad[mask] = true
	}

	for bit := range m {
		for state := range limit {
			if (state>>bit)&1 == 1 && bad[state^(1<<bit)] {
				bad[state] = true
			}
		}
	}

	best := 0
	for state := range limit {
		if !bad[state] {
			best = max(best, digitCount(state))
		}
	}

	return m - best
}

func digitCount(num int) int {
	return bits.OnesCount(uint(num))
}
