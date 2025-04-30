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
	s := `3 2 0 2
0 1 3
2 5
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 11 10 20
6 18 32 63 66 68 87
6 8 15 23 25 41 53 59 60 75 90
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20 18 1 2
-9999944 -9999861 -9999850 -9999763 -9999656 -9999517 -9999375 -9999275 -9999203 -9999080 -9998988 -9998887 -9998714 -9998534 -9998475 -9998352 -9998164 -9998016 -9998002 -9997882
-9999976 -9999912 -9999788 -9999738 -9999574 -9999460 -9999290 -9999260 -9999146 -9999014 -9998962 -9998812 -9998616 -9998452 -9998252 -9998076 -9997928 -9997836
`
	expect := 2
	runSample(t, s, expect)
}
