package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for i, v := range expect {
		if res[i] != v {
			t.Errorf("Sample expect %v, but got %v, differ at %d", expect, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
67 0
6 1
7 1
1 0
100 0
62 1
`
	expect := []int{67, 100, 69}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1
2 2
4 2
3 1
`
	expect := []int{7}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 2
6 1
7 0
8 1
`
	expect := []int{7, 14}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 10
643444063 7
208999612 5
148423326 0
543584678 3
330940797 7
155699366 3
384886181 3
382481223 3
756968649 1
6107718 4
939210119 3
409114343 0
706144299 1
418198997 2
959553978 2
968610514 5
990506063 2
817619952 3
`
	expect := []int{1954396145, 1954396145, 2511125041, 1954396145, 1954396145, 1954396145, 2146582719, 2540525436, 2177534804, 2389534874}
	runSample(t, s, expect)
}
