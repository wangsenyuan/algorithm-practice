package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	a := readNNums(reader, len(res))
	n := len(res)

	for i := k; i < n; i++ {
		if res[i-k] >= res[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	abs_sum := func(arr []int) int {
		var sum int
		for _, v := range arr {
			sum += abs(v)
		}
		return sum
	}
	if abs_sum(res) != abs_sum(a) {
		t.Fatalf("Sample expect %v, but got %v", a, res)
	}
}

func abs(num int) int {
	return max(num, -num)
}
func TestSample1(t *testing.T) {
	s := `3 2
? 1 2
0 1 2`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `5 1
-10 -9 ? -7 -6
-10 -9 -8 -7 -6 
`
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := `5 3
4 6 7 2 9
`
	runSample(t, s, false)
}

func TestSample4(t *testing.T) {
	s := `5 2
-1 ? ? ? 1
-1 -1 0 0 1
`
	runSample(t, s, true)
}
