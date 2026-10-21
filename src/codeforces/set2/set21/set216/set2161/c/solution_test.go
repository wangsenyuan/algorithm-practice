package main

import (
	"bufio"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	x, points, order := drive(reader)

	play := func(arr []int) int {
		var sum, score int
		for _, v := range arr {
			sum += v
			if sum >= x {
				score += v
			}
			sum %= x
		}
		return score
	}

	u := play(expect)
	v := play(order)
	if u != v || v != points {
		t.Fatalf("Sample expect %d %v, but got %d %v", u, expect, points, order)
	}
	slices.Sort(order)
	slices.Sort(expect)
	if !reflect.DeepEqual(order, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, order)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 2
1 2 1 2 1 2 1 2 1 2
`, []int{1, 2, 2, 2, 2, 2, 1, 1, 1, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 10
2 2 2 2 5
`, []int{2, 2, 2, 2, 5})
}

func TestSample3(t *testing.T) {
	runSample(t, `11 23
5 5 22 1 21 2 10 3 1 1 2
`, []int{1, 1, 5, 2, 1, 2, 5, 3, 10, 21, 22})
}

func TestSample4(t *testing.T) {
	runSample(t, `1 1
1
`, []int{1})
}

func TestSample5(t *testing.T) {
	runSample(t, `1 17
11
`, []int{11})
}

func TestSample6(t *testing.T) {
	runSample(t, `3 100
44 32 1
`, []int{44, 32, 1})
}

func TestSample7(t *testing.T) {
	runSample(t, `16 100500
42801 73112 95296 68791 42217 21871 29316 84405 24273 42894 63370 53473 57156 61369 80 27290
`, []int{53473, 42894, 80, 57156, 42801, 61369, 42217, 63370, 29316, 68791, 27290, 73112, 24273, 84405, 21871, 95296})
}
