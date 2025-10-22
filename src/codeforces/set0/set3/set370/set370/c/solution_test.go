package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	w, res := drive(reader)
	if w != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, w)
	}

	var cnt int
	for _, cur := range res {
		if cur[0] != cur[1] {
			cnt++
		}
	}

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}
}

func TestSample1(t *testing.T) {
	s := `6 3
1 3 2 2 1 1
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 2 1 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 6
4 3 4 4 4 4 4 3 4 4 4 1 2 4 2 6 4 4 3 2 4 4 4 4 3 4 4 2 4 4 4 6 4 1 4 2 4 4 4 4 4 4 4 4 6 6 4 4 4 4 4 1 4 5 4 4 4 4 4 4 4 4 4 4 4 4 2 4 4 4 4 4 4 4 5 4 2 4 4 4 3 4 5 4 6 4 5 4 4 4 2 4 4 6 4 3 4 5 3 4
`
	expect := 58
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 3
1 1 2 2 3 3`
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 2
1 2 1 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `5 3
1 2 3 1 2`
// 1 3
// 2 1
// 3 2
// 1 
	expect := 5
	runSample(t, s, expect)
}
