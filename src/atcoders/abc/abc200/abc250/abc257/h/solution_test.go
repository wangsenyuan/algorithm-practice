package main

import (
	"bufio"
	"math/rand"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
1 2 3
1 1 1 1 1 1
2 2 2 2 2 2
3 3 3 3 3 3
`, 20)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 5
2 5 6 5 2 1 7 9 7 2
5 5 2 4 7 6
2 2 8 7 7 9
8 1 9 6 10 8
8 6 10 3 3 9
1 10 5 8 1 10
7 8 4 8 6 5
1 10 2 5 1 7
7 4 1 4 5 4
5 10 1 5 1 2
5 1 2 3 6 2
`, 1014)
}

func bruteSolve(n, k int, c []int, a [][]int) int {
	const bruteMod = 998244353
	best := int64(-1 << 60)
	var dfs func(int, int, int64, int64)
	dfs = func(i, cnt int, sumX, sumY int64) {
		if cnt == k {
			cur := sumX*sumX + sumY
			if cur > best {
				best = cur
			}
			return
		}
		if i == n || cnt+n-i < k {
			return
		}
		dfs(i+1, cnt, sumX, sumY)

		var x, q int64
		for _, v := range a[i] {
			x += int64(v)
			q += int64(v * v)
		}
		y := 6*q - x*x - 36*int64(c[i])
		dfs(i+1, cnt+1, sumX+x, sumY+y)
	}
	dfs(0, 0, 0, 0)
	return int((best%bruteMod + bruteMod) % bruteMod * brutePow(36, bruteMod-2) % bruteMod)
}

func brutePow(a, b int64) int64 {
	const bruteMod = 998244353
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % bruteMod
		}
		a = a * a % bruteMod
		b >>= 1
	}
	return res
}

func TestAgainstBruteForceSmall(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for tc := 0; tc < 200; tc++ {
		n := 1 + rng.Intn(8)
		k := 1 + rng.Intn(n)
		c := make([]int, n)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			c[i] = 1 + rng.Intn(20)
			a[i] = make([]int, 6)
			for j := 0; j < 6; j++ {
				a[i][j] = 1 + rng.Intn(20)
			}
		}
		expect := bruteSolve(n, k, c, a)
		if res := solve(n, k, c, a); res != expect {
			t.Fatalf("case %d expect %d, but got %d (n=%d k=%d c=%v a=%v)", tc, expect, res, n, k, c, a)
		}
	}
}
