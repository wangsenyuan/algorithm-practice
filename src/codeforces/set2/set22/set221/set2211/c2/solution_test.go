package main

import (
	"bufio"
	"math/rand"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
1 2 3 4 5
3 1 5 2 4
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
1 2 1 2 1
2 -1 -1 -1 -1
`, "YES")
}

func TestSample3(t *testing.T) {
	runSample(t, `6 1
5 6 2 2 4 3
5 -1 -1 2 -1 3
`, "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, `2 1
1 2
2 -1
`, "NO")
}

func TestSample5(t *testing.T) {
	runSample(t, `6 4
1 2 3 4 1 2
2 -1 3 -1 4 -1
`, "NO")
}

func TestFixedValueMustMatchInitialWindow(t *testing.T) {
	runSample(t, `2 1
2 2
-1 1
`, "NO")
}

func valid(k int, a []int, b []int) bool {
	n := len(a)
	for l := 0; l+k <= n; l++ {
		cnt := make(map[int]int)
		for i := l; i < l+k; i++ {
			cnt[a[i]]++
			cnt[b[i]]--
		}
		for _, v := range cnt {
			if v != 0 {
				return false
			}
		}
	}
	return true
}

func bruteForce(k int, a []int, b []int) string {
	n := len(a)
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == n {
			return valid(k, a, b)
		}
		if b[pos] != -1 {
			return dfs(pos + 1)
		}
		for x := 1; x <= n; x++ {
			b[pos] = x
			if dfs(pos + 1) {
				b[pos] = -1
				return true
			}
		}
		b[pos] = -1
		return false
	}
	if dfs(0) {
		return "YES"
	}
	return "NO"
}

func TestSolveAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for n := 1; n <= 8; n++ {
		for k := 1; k <= n; k++ {
			for it := 0; it < 100; it++ {
				a := make([]int, n)
				b := make([]int, n)
				for i := range n {
					a[i] = rng.Intn(n) + 1
					b[i] = rng.Intn(n) + 1
					if rng.Intn(3) == 0 {
						b[i] = -1
					}
				}
				b1 := append([]int(nil), b...)
				b2 := append([]int(nil), b...)
				expect := bruteForce(k, a, b1)
				if res := solve(k, a, b2); res != expect {
					t.Fatalf("n=%d k=%d a=%v b=%v: got %s, want %s", n, k, a, b, res, expect)
				}
			}
		}
	}
}
