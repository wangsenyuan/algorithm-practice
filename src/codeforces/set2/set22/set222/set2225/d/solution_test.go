package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `8 1
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `15 8
`, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, `10 10
`, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, `5989566119 1996588700
`, 99996)
}

func TestSolveAgainstBruteForce(t *testing.T) {
	for n := int64(1); n <= 80; n++ {
		for x := int64(1); x <= n; x++ {
			expect := bruteForce(n, x)
			res := solve(n, x)
			if res != expect {
				t.Fatalf("n=%d x=%d, expect %d, got %d", n, x, expect, res)
			}
		}
	}
}

func TestLargeBounds(t *testing.T) {
	tests := []struct {
		n int64
		x int64
	}{
		{1_000_000_000_000_000_000, 1},
		{1_000_000_000_000_000_000, 500_000_000_000_000_000},
		{1_000_000_000_000_000_000, 999_999_999_999_999_999},
		{1_000_000_000_000_000_000, 1_000_000_000_000_000_000},
	}

	for _, cur := range tests {
		res := solve(cur.n, cur.x)
		if res < 0 || res >= mod {
			t.Fatalf("n=%d x=%d, got out-of-range result %d", cur.n, cur.x, res)
		}
	}
}

func bruteForce(n, x int64) int64 {
	var ans int64
	for l := int64(1); l <= x; l++ {
		var xor int64
		for r := l; r <= n; r++ {
			xor ^= r
			if r >= x && xor == 0 {
				ans++
			}
		}
	}
	return ans
}
