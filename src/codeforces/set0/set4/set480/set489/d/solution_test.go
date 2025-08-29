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
	s := `5 4
1 2
2 3
1 4
4 3
	`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 12
1 2
1 3
1 4
2 1
2 3
2 4
3 1
3 2
3 4
4 1
4 2
4 3
	`
	expect := 12
	runSample(t, s, expect)
}
