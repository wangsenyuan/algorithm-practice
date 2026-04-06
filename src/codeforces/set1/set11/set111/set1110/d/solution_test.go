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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 6
2 3 3 3 4 4 4 5 5 6
`
	// 2 3 4
	// 3 4 5
	// 4 5 6
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `12 6
1 5 3 3 3 4 3 5 3 2 3 3
`
	expect := 3
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `13 5
1 1 5 1 2 3 3 2 4 2 3 4 5
`
	expect := 4
	runSample(t, s, expect)
}
