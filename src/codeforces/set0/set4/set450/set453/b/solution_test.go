package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)

	get := func(b []int) int {
		var res int
		for i, x := range a {
			res += abs(x - b[i])
		}
		return res
	}
	s1 := get(expect)
	s2 := get(res)

	if s1 != s2 {
		t.Fatalf("Sample result %v, not getting minimum sum %d, but got %d", res, s1, s2)
	}

	for i := range len(res) {
		for j := range i {
			if gcd(res[i], res[j]) != 1 {
				t.Fatalf("Sample result %v, not getting coprime pairs", res)
			}
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	s := `5
1 1 1 1 1
`
	expect := []int{1, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 6 4 2 8
`
	expect := []int{1, 5, 3, 1, 8}
	runSample(t, s, expect)
}

// 这个case跑起来太慢了
// func TestSample3(t *testing.T) {
// 	s := `10
// 16 3 16 10 12 5 14 14 15 27
// `
// 	expect := []int{19, 1, 17, 7, 11, 1, 16, 13, 15, 29}
// 	runSample(t, s, expect)
// }
