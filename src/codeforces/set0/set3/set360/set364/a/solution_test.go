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
	s := `10
12345
`
	runSample(t, s, 6)
}

func TestSample2(t *testing.T) {
	s := `16
439873893693495623498263984765
`
	runSample(t, s, 40)
}

func TestSample3(t *testing.T) {
	s := `0
1230
`

//   1 2 3 0
// 1       0
// 2       0
// 3       0
// 0 0 0 0 0

	runSample(t, s, 19)
}
