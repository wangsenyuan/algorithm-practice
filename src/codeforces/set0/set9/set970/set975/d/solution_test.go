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
	s := `4 1 1
1 -1 -1
2 1 1
3 1 1
4 -1 -1
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 0
-1 1 0
0 0 -1
1 -1 -2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1 0
0 0 0
1 0 0
2 0 0
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 7 -626288749
795312099 49439844 266151109
-842143911 23740808 624973405
-513221420 -44452680 -391096559
-350963348 -5068756 -160670209
690883790 11897718 3356227
-509035268 -45646185 -210137445
-121282138 -32581578 230716703
491731655 9500548 -13423963
-665038289 48170248 446577586
495114076 -38468595 -159894315
`
	expect := 20
	runSample(t, s, expect)
}
