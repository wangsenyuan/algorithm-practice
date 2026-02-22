package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, a, res := drive(reader)

	if len(res) != k {
		t.Fatalf("Sample expect %d, but got %d", k, len(res))
	}
	marked := make(map[int]bool)
	var arr []int
	for _, cur := range res {
		if len(cur) == 0 {
			t.Fatalf("Sample result %v, is not correct", res)
		}
		var tmp int
		for _, v := range cur {
			tmp += v
		}
		if marked[tmp] {
			t.Fatalf("Sample result %v, is not correct", res)
		}
		marked[tmp] = true
		arr = append(arr, cur...)
	}

	slices.Sort(arr)
	slices.Sort(a)
	var i int
	for _, v := range arr {
		for i < len(a) && a[i] < v {
			i++
		}
		if i == len(a) || a[i] != v {
			t.Fatalf("Sample result %v, is not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 6
1 2 3
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	// Random test: 10 distinct positive numbers, k = 55
	s := `10 55
7 12 3 19 5 22 8 14 31 6
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `50 836
43 33 24 13 29 34 11 17 39 14 40 23 35 26 38 28 8 32 4 25 46 9 5 21 45 7 6 30 37 12 2 10 3 41 42 22 50 1 18 49 48 44 47 19 15 36 20 31 16 27
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `50 1260
4 20 37 50 46 19 25 47 10 6 34 12 41 9 22 28 40 42 15 27 8 38 17 13 7 30 48 23 11 16 2 32 18 24 14 33 49 35 44 39 3 36 31 45 1 29 5 43 26 21
`
	runSample(t, s)
}
