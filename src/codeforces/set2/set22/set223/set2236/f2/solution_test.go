package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
2 2
2 4
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
1 5
5
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
7 4
2 4 8 13 111 6 7
`, 360)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
3 1000
1 2 3
`, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
3 3
4 8 10
`, 0)
}

func TestAgainstBruteForceSmall(t *testing.T) {
	samples := []struct {
		a []int
		x int
	}{
		{[]int{1}, 1},
		{[]int{2, 4}, 2},
		{[]int{2, 3, 6}, 1},
		{[]int{4, 8, 10}, 3},
		{[]int{6, 10, 15}, 6},
		{[]int{8, 9, 12, 5}, 4},
		{[]int{12, 18, 20}, 12},
	}
	for _, cur := range samples {
		expect := bruteForce(cur.a, cur.x)
		if res := solve(cur.a, cur.x); res != expect {
			t.Fatalf("solve(%v, %d) = %d, expect %d", cur.a, cur.x, res, expect)
		}
	}
}

func bruteForce(a []int, x int) int {
	divs := make([][]int, len(a))
	for i, v := range a {
		for d := 1; d <= v; d++ {
			if v%d == 0 {
				divs[i] = append(divs[i], d)
			}
		}
	}
	var dfs func(int, int64, int64) int
	dfs = func(i int, prod int64, l int64) int {
		if i == len(a) {
			if int64(x)*l == prod {
				return 1
			}
			return 0
		}
		var res int
		for _, d := range divs[i] {
			res += dfs(i+1, prod*int64(d), lcm(l, int64(d)))
		}
		return res
	}
	return dfs(0, 1, 1)
}

func lcm(a, b int64) int64 {
	return a / gcd64(a, b) * b
}

func gcd64(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
