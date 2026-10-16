package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, _, res := drive(reader)
	for _, line := range res {
		if len(line) == 0 {
			fmt.Fprintln(writer, "NIE")
		} else {
			fmt.Fprintln(writer, line[0], line[1])
		}
	}
}

func drive(reader *bufio.Reader) (s string, queries []int, res [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	fmt.Fscan(reader, &s)
	s = s[:n]
	queries = make([]int, m)
	for i := range queries {
		fmt.Fscan(reader, &queries[i])
	}
	res = solve(s, queries)
	return
}

func solve(s string, queries []int) [][]int {
	n := len(s)
	ansL := make([]int32, 2*n+1)
	ansR := make([]int32, 2*n+1)
	best := [2]int32{-1, -1}

	firstPos := [2]int32{-1, -1}
	firstPos[0] = 0

	pref := make([]int32, n+1)
	for i := range n {
		pref[i+1] = pref[i] + 1
		if s[i] == 'T' {
			pref[i+1]++
		}
		for parity := range 2 {
			start := firstPos[parity]
			if start < 0 || int(start) > i {
				continue
			}
			cur := pref[i+1] - pref[start]
			if best[cur%2] < cur {
				best[cur%2] = cur
				ansL[int(cur)] = start
				ansR[int(cur)] = int32(i)
			}
		}
		w := pref[i+1] % 2
		if firstPos[w] < 0 {
			firstPos[w] = int32(i + 1)
		}
	}

	play := func(sum int32, l int32, r int32) {
		for l <= r && sum >= 0 {
			ansL[int(sum)] = l + 1
			ansR[int(sum)] = r + 1
			if s[l] == 'T' {
				l++
			} else if s[r] == 'T' {
				r--
			} else {
				l++
				r--
			}
			sum -= 2
		}
	}

	if best[0] >= 0 {
		play(best[0], ansL[best[0]], ansR[best[0]])
	}
	if best[1] >= 0 {
		play(best[1], ansL[best[1]], ansR[best[1]])
	}

	res := make([][]int, len(queries))
	for i, x := range queries {
		if x <= 2*n && ansL[x] > 0 {
			res[i] = []int{int(ansL[x]), int(ansR[x])}
		}
	}
	return res
}
