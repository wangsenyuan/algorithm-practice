package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, n, files := process(reader)
	if len(res) > 2*n {
		t.Fatalf("Sample expect at most %d moves, but got %v", 2*n, res)
	}
	belong := make([]int, n+1)
	for i := range n + 1 {
		belong[i] = -1
	}
	var tot int
	for i, file := range files {
		for _, c := range file {
			belong[c] = i
		}
		tot += len(file)
	}

	for _, move := range res {
		x, y := move[0], move[1]
		belong[x], belong[y] = belong[y], belong[x]
	}

	marked := make([]bool, len(files))
	for pos := 1; pos < tot; {
		id := belong[pos]
		if id < 0 || marked[id] {
			t.Fatalf("Sample result %v, not correct", belong)
		}
		marked[id] = true

		sz := len(files[id])

		for i := pos; i < pos+sz; i++ {
			if belong[i] != id {
				t.Fatalf("Sample result %v, not correct", belong)
			}
		}

		pos += sz
	}
}

func TestSample1(t *testing.T) {
	s := `7 2
2 1 2
3 3 4 5
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `7 2
2 1 3
3 2 4 5
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `7 2
2 1 7
3 2 4 5
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `7 3
2 1 6
3 2 4 5
1 7
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `3 1
2 3 1
`
	runSample(t, s)
}
