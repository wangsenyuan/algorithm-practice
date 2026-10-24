package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	n, dances := parseCase(s)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !validColoring(n, dances, res) {
		t.Fatalf("invalid coloring %v for dances %v", res, dances)
	}
}

func parseCase(s string) (int, [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))
	var n, m int
	fmt.Fscan(reader, &n, &m)
	dances := make([][]int, m)
	for i := range m {
		dances[i] = make([]int, 3)
		fmt.Fscan(reader, &dances[i][0], &dances[i][1], &dances[i][2])
	}
	return n, dances
}

func validColoring(n int, dances [][]int, colors []int) bool {
	if len(colors) != n {
		return false
	}
	for _, color := range colors {
		if color < 1 || color > 3 {
			return false
		}
	}
	for _, dance := range dances {
		seen := 0
		for _, dancer := range dance {
			if dancer < 1 || dancer > n {
				return false
			}
			bit := 1 << colors[dancer-1]
			if seen&bit != 0 {
				return false
			}
			seen |= bit
		}
		if seen != 1<<1|1<<2|1<<3 {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, `7 3
1 2 3
1 4 5
4 6 7
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `9 3
3 6 9
2 5 8
1 4 7
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 2
4 1 5
3 1 2
`)
}
