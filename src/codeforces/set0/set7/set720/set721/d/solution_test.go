package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, k, x, b := drive(reader)

	prod := func(arr []int) int {
		res := 1
		for _, num := range arr {
			res *= num
		}
		return res
	}

	w := prod(expect)
	v := prod(b)

	if w != v {
		t.Fatalf("Sample expect %v, but got %v", expect, b)
	}

	mod := func(num int) int {
		num %= x
		if num < 0 {
			num += x
		}
		return num
	}

	var cnt int

	for i, v := range a {
		w := b[i]

		if mod(w) != mod(v) {
			t.Fatalf("Sample result %v, not correct at position %d", b, i)
		}
		c := max(v, w) - min(v, w)
		cnt += c / x
	}

	if cnt > k {
		t.Fatalf("Sample result %v, used %d operations, but expected %d", b, cnt, k)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3 1
5 4 3 5 2
`
	expect := []int{5, 4, 3, 5, -1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3 1
5 4 3 5 5
`
	expect := []int{5, 4, 0, 5, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 3 1
5 4 4 5 5
`
	expect := []int{5, 1, 4, 5, 5}
	// 5 * 1 * 4 * 5 * 5 = 500
	// 4 * 4 * 4 * 4 * 4 = 1024
	// a * b 如果 a < b
	// (a - x) * b = a * b - b * x
	// a * (b - x) = a * b - a * x
	// 显然减去最小的那个更好
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 2 7
5 4 2
`
	expect := []int{5, 11, -5}

	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 1 3
-5 -4 6
`

	// 5 * 4 * 6 = 120
	// 8 * 4 * 6 = 212
	//
	expect := []int{-5, -1, 6}

	runSample(t, s, expect)
}
