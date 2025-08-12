package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2
1 3
1 -1 1`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 3
4 5
2 5
1 3
0 2 1 4 3`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12
1 6
10 1
4 1
7 1
1 2
5 1
1 8
1 11
3 1
12 1
9 1
580660007 861441526 -264928594 488291045 253254575 -974301934 709266786 926718320 87511873 514836444 -702876508 848928657`
	expect := 2529263875
	runSample(t, s, expect)
}
