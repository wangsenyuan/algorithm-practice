package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 1
2 0 4 0 6 0`
	expect := []int{2, 4, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 2
2 3 0 5 6 0`
	expect := []int{2, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
0 0 0 0`
	expect := []int{1, 2, 3, 4}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 2
0 0 1 0 4 5`
	expect := []int{1, 3, 4, 6}
	runSample(t, s, expect)
}
