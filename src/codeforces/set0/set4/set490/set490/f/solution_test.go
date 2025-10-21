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
	s := `6
1 2 3 4 5 1
1 2
2 3
3 4
3 5
3 6
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 5
1 2
1 3
2 4
3 5
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
73717 73738 73675 73696
3 4
1 2
4 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
786 764 797 775
3 1
3 2
3 4
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2
1 2
1 2
`
	expect := 2
	runSample(t, s, expect)
}
