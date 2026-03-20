package main

import "testing"

func verifySchedule(t *testing.T, n int, res [][]int) {
	t.Helper()

	seen := make([][]bool, n)
	for i := 0; i < n; i++ {
		seen[i] = make([]bool, n)
	}

	for _, team := range res {
		mark := make([]bool, n)
		for _, x := range team {
			if x < 1 || x > n {
				t.Fatalf("player %d out of range for n=%d", x, n)
			}
			if mark[x-1] {
				t.Fatalf("duplicate player %d in one practice", x)
			}
			mark[x-1] = true
		}
		if len(team) == 0 || len(team) == n {
			t.Fatalf("invalid team size %d for n=%d", len(team), n)
		}

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if mark[i] != mark[j] {
					seen[i][j] = true
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !seen[i][j] {
				t.Fatalf("pair (%d, %d) was never separated", i+1, j+1)
			}
		}
	}
}

func ceilLog2(n int) int {
	var ans int
	for x := 1; x < n; x <<= 1 {
		ans++
	}
	return ans
}

func runSample(t *testing.T, n int) {
	t.Helper()
	res := solve(n)
	verifySchedule(t, n, res)

	expect := ceilLog2(n)
	if len(res) != expect {
		t.Fatalf("expect %d practices, but got %d", expect, len(res))
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3)
}

func TestPowerOfTwo(t *testing.T) {
	runSample(t, 8)
}

func TestNonPowerOfTwo(t *testing.T) {
	runSample(t, 10)
}
