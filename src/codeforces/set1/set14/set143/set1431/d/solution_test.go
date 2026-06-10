package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, ord := drive(reader)

	play := func(res []int) int {
		var cnt int
		var marker int
		for _, i := range res {
			i--
			if a[i] >= marker {
				cnt++
				marker = 0
			}
			marker++
		}
		return cnt + 1
	}
	x := play(expect)
	y := play(ord)

	if x != y {
		t.Fatalf("Sample result %v, is not correct, as it uses %d, instead of %d", ord, y, x)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2 1 2`
	expect := []int{4, 1, 3, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
2 1`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 1 1`
	expect := []int{3, 1, 2}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
2 3 1 3`
	expect := []int{4, 3, 2, 1}
	runSample(t, s, expect)
}
