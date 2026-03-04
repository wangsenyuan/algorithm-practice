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
	s := "10 3 2"
	// L = 1, 6, 7

	runSample(t, s, []int{3, 10})
}

func TestSample2(t *testing.T) {
	s := "7 1 2"

	runSample(t, s, []int{3, 7})
}

func TestSample3(t *testing.T) {
	s := "3000000000000000000 2999999999999999873 2999999999999999977"
	runSample(t, s, []int{23437499999999999, 23437500000000000})
}

func TestSample4(t *testing.T) {
	s := "1000000000000000000 1000000000 2000000000"
	runSample(t, s, []int{1, 2})
}
