package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res, X, Y, Z := drive(reader)

	if (len(expect) == 0) != (len(res) == 0) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if len(expect) == 0 {
		return
	}

	play := func(a []int) int {
		for i := range len(X) {
			x := X[i] - 1
			y := Y[i] - 1
			if a[x]^a[y] != Z[i] {
				t.Fatalf("Sample result %v is not valid", a)
			}
		}
		var sum int
		for _, v := range a {
			sum += v
		}
		return sum
	}

	sum1 := play(res)
	sum2 := play(expect)
	if sum1 != sum2 {
		t.Fatalf("Sample result %v is not correct got %d, expect %d", res, sum1, sum2)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
1 3 4
1 2 3
`, []int{0, 3, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 3 4
1 2 3
2 3 5
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 8
4 2 4
2 3 11
3 4 15
4 5 6
3 2 11
3 3 0
3 1 9
3 4 15
`, []int{0, 2, 9, 6, 0})
}
