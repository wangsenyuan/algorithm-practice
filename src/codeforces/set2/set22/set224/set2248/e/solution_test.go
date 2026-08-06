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
	runSample(t, `6 4 3
2 5
3 9
4 1
5 3
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `7 3 5
2 5
4 5
7 10
`, "NO")
}

func TestSample3(t *testing.T) {
	runSample(t, `684492057 3 386217943
367971233 991739271
612599954 429216213
684492056 402931836
`, "YES")
}

func value(a []int, n, d int64, p []int64, r []int64) int64 {
	var res int64
	var cnt int64
	for _, bit := range a {
		if bit == 1 {
			res += d
			cnt++
		} else {
			cnt = 0
		}
		for i, x := range p {
			if cnt == x {
				res += r[i]
				break
			}
		}
		if cnt == n {
			cnt = 0
		}
	}
	return res
}

func bruteForce(n, d int64, p []int64, r []int64) string {
	for length := 1; length <= int(2*n+1); length++ {
		ones := make([]int, length)
		for i := range ones {
			ones[i] = 1
		}
		base := value(ones, n, d, p, r)
		for mask := 0; mask < 1<<length; mask++ {
			a := make([]int, length)
			for i := range a {
				a[i] = mask >> i & 1
			}
			if value(a, n, d, p, r) > base {
				return "YES"
			}
		}
	}
	return "NO"
}

func TestSolveAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for n := int64(1); n <= 6; n++ {
		for it := 0; it < 200; it++ {
			d := int64(rng.Intn(6))
			var p []int64
			var r []int64
			for x := int64(1); x <= n; x++ {
				if rng.Intn(2) == 0 {
					p = append(p, x)
					r = append(r, int64(rng.Intn(8)+1))
				}
			}
			expect := bruteForce(n, d, p, r)
			if res := solve(n, d, p, r); res != expect {
				t.Fatalf("n=%d d=%d p=%v r=%v: got %s, want %s", n, d, p, r, res, expect)
			}
		}
	}
}
