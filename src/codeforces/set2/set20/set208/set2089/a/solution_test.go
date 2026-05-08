package main

import "testing"

func runSample(t *testing.T, n int) {
	res := solve(n)

	marked := make([]bool, n+1)
	var sum int
	var cnt int
	for i, v := range res {
		marked[v] = true
		sum += v
		avg := (sum + i) / (i + 1)
		if checkPrime(avg) {
			cnt++
		}
	}
	if cnt < n/3-1 {
		t.Fatalf("Sample result %v, is invalid, as it has only %d prime averages", res, cnt)
	}
	for i := 1; i <= n; i++ {
		if !marked[i] {
			t.Fatalf("Sample result %v, is invalid, as it does not contain %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 100)
}

func TestSample3(t *testing.T) {
	runSample(t, 107)
}