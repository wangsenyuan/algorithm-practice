package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, m, k := process(reader)
	expect := readNum(reader)

	f := func(arr []int) int {
		rem := arr[m*k:]
		freq := make(map[int]int)
		var mex int
		for _, x := range rem {
			freq[x]++
			for freq[mex] > 0 {
				mex++
			}
		}
		return mex
	}
	x := f(res)
	if x != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, x)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 1 1
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2 2
1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `6 1 4
2`)
}

func TestSample4(t *testing.T) {
	runSample(t, `8 2 2
2`)
}

func TestSample5(t *testing.T) {
	runSample(t, `8 1 5
3`)
}

func TestSample6(t *testing.T) {
	runSample(t, `11 3 3
2`)
}