package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3 2
1 30
1 40
1 50
2 10
3 20
`
	runSample(t, s, 110)
}

func TestSample2(t *testing.T) {
	s := `5 3 3
1 30
1 40
1 50
2 10
3 20
`
	runSample(t, s, 80)
}

func TestSample3(t *testing.T) {
	s := `5 5 1
4 1000000000
5 1000000000
4 1000000000
5 1000000000
4 1000000000
`
	runSample(t, s, 5000000000)
}
