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
	s := `1 1
2 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
1 3
2 2
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3
3 10
2 7
2 8
1 1
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 18
3 3
1 10
2 10
3 6
1 3
2 3
`
	expect := 35
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 2
3 4
3 5
3 7
1 10
1 2
1 2
1 8
3 2
1 8
3 3
`
	expect := 18
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6 6
1 6
1 4
1 8
3 2
3 2
2 8
`
	expect := 26
	runSample(t, s, expect)
}
