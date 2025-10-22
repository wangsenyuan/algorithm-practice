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
	s := `5 7
0 1 0 2 0
0 0 4 10 0
0 0 0 0 5
0 0 0 0 10
0 0 0 0 0
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 10
0 1 0 0 0
0 0 2 0 0
0 0 0 3 0
0 0 0 0 4
100 0 0 0 0
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 930
0 994194 999113 0 0 0 0 0
0 0 1 991020 0 0 0 0
0 1 0 0 992293 0 0 0
0 0 0 0 1 996760 0 0
0 0 0 1 0 0 991401 0
0 0 0 0 0 0 1 991401
0 0 0 0 0 1 0 998645
0 0 0 0 0 0 0 0
`
	expect := 1983352
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 0
0 1 1 1 1 1 1 1 1 1
1 0 1 1 1 1 1 1 1 1
1 1 0 1 1 1 1 1 1 1
1 1 1 0 1 1 1 1 1 1
1 1 1 1 0 1 1 1 1 1
1 1 1 1 1 0 1 1 1 1
1 1 1 1 1 1 0 1 1 1
1 1 1 1 1 1 1 0 1 1
1 1 1 1 1 1 1 1 0 1
1 1 1 1 1 1 1 1 1 0
`
	expect := 9
	runSample(t, s, expect)
}
