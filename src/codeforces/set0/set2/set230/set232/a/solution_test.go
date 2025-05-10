package main

import "testing"

func runSample(t *testing.T, k int) {
	res := solve(k)
	if len(res) > 100 {
		t.Fatalf("Sample result is too long having %d nodes", len(res))
	}
	n := len(res)

	var cnt int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if res[i][j] == '0' {
				continue
			}
			for u := j + 1; u < n; u++ {
				if res[i][u] == '1' && res[j][u] == '1' {
					cnt++
				}
			}
		}
	}

	if cnt != k {
		t.Fatalf("Sample result is not correct, expect %d, but got %d", k, cnt)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 29257)
}
