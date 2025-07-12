package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	var expect int
	fmt.Fscan(reader, &expect)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
2 1 2
2 2 3
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 3
3 1 2 3
6`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 4
2 1 2
3 2 3 4
2`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 2
3 2 2 1
2 1 2
1`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `3 7
2 1 2
2 3 4
3 5 6 7
24`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `10 100
16 25 7 48 43 16 23 66 3 17 31 64 27 7 17 11 60
62 76 82 99 77 19 26 66 46 9 54 77 8 34 76 70 48 53 35 69 29 84 22 16 53 36 27 24 81 2 86 67 45 22 54 96 37 8 3 22 9 30 63 61 86 19 16 47 3 72 39 36 1 50 1 18 7 44 52 66 90 3 63
3 22 61 39
9 28 69 91 62 98 23 45 9 10
2 42 20
3 90 46 55
2 71 9
1 7
1 44
1 94
732842622`
	runSample(t, s)
}
