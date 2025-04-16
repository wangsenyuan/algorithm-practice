package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 2
2 1
3 4
4 3
	`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
1 4
2 5
	`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 5
1 2
2 3
3 1
2 4
4 1
	`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 4
1 1
1 2
2 1
1 2
	`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 3
1 3
1 2
2 3
	`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4 4
1 2
2 3
3 4
1 4
	`
	expect := 2
	runSample(t, s, expect)
}
