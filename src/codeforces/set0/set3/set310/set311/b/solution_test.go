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

func TestNegativeDepartureIsAllowed(t *testing.T) {
	s := `2 1 1
3
2 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestAllCatsCanBePickedWithoutWaiting(t *testing.T) {
	s := `2 2 1
10
2 1
2 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample1(t *testing.T) {
	s := `4 6 2
1 3 5
1 0
2 1
4 9
1 10
2 10
3 12
`
	expect := 3
	runSample(t, s, expect)
}

func TestTwoKeepersSplitSortedCats(t *testing.T) {
	s := `3 4 2
5 5
3 12
3 13
3 21
3 22
`
	expect := 2
	runSample(t, s, expect)
}
