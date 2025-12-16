package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
koyomi
3
1 o
4 o
4 m
`
	expect := []int{3, 6, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `15
yamatonadeshiko
10
1 a
2 a
3 a
4 a
5 a
1 b
2 b
3 b
4 b
5 b
`
	expect := []int{3,
		4,
		5,
		7,
		8,
		1,
		2,
		3,
		4,
		5,
	}
	runSample(t, s, expect)
}
