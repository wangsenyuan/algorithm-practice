package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)
	n := len(a)
	check := func(arr []int) int {
		var res int
		for i := range arr {
			if arr[i] > n {
				t.Fatalf("Sample result %v, not valid", arr)
			}
			l, r := i, i
			for l >= 0 && r < len(arr) && arr[l] == arr[r] {
				l--
				r++
			}
			res += r - i
			l, r = i, i+1
			for l >= 0 && r < len(arr) && arr[l] == arr[r] {
				l--
				r++
			}
			res += i - l
		}
		return res
	}

	x := check(append(a, expect...))
	y := check(append(a, res...))

	if x != y {
		t.Fatalf("Sample expect %v(%d), but got %v(%d)", expect, x, res, y)
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
1 3 3 4`
	expect := []int{2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
2 2 2 2`
	expect := []int{1, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 1
4 1 5 2 2`
	expect := []int{3}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 3
1 2 3 4 5 6`
	expect := []int{3, 4, 1}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 3
3 2 5 2 3`
	expect := []int{4, 1, 5}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 3
1 2 2`
	expect := []int{3, 1, 2}
	runSample(t, s, expect)
}
