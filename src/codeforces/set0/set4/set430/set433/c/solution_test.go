package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 6
1 2 3 4 3 2
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 5
9 4 3 8 8
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 10
2 5 2 2 3 5 3 2 1 3
`

// 如果把 x merge 到 y， 那么所有的 a[?] = x 都需要变成 a[?] = y
	expect := 7
	runSample(t, s, expect)
}
