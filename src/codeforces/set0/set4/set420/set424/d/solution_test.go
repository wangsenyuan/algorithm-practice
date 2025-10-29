package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	t1, tp, tu, td, a, res := drive(reader)

	move := func(a int, b int) int {
		if a < b {
			return tu
		}
		if a > b {
			return td
		}
		return tp
	}

	play := func(r1, c1, r2, c2 int) int {
		r1--
		r2--
		c1--
		c2--
		var res int
		for j := c1 + 1; j <= c2; j++ {
			res += move(a[r1][j-1], a[r1][j])
		}
		for i := r1 + 1; i <= r2; i++ {
			res += move(a[i-1][c2], a[i][c2])
		}
		for j := c2 - 1; j >= c1; j-- {
			res += move(a[r2][j+1], a[r2][j])
		}
		for i := r2 - 1; i >= r1; i-- {
			res += move(a[i+1][c1], a[i][c1])
		}
		return res
	}

	u := play(expect[0], expect[1], expect[2], expect[3])
	v := play(res[0], res[1], res[2], res[3])

	if abs(u-t1) != abs(v-t1) {
		t.Fatalf("Sample expect %v(%d), but got %v(%d)", expect, u, res, v)
	}
}

func TestSample1(t *testing.T) {
	s := `6 7 48
3 6 2
5 4 8 3 3 7 9
4 1 6 8 7 1 1
1 6 4 6 4 8 6
7 2 6 1 6 9 4
1 9 8 6 3 9 2
4 5 6 8 4 3 7`
	expect := []int{1, 1, 4, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 10
1 3 2
1 2 3
3 4 5
5 6 7`
	expect := []int{1, 1, 3, 3}
	runSample(t, s, expect)
}
