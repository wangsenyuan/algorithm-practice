package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	// 1, 2, 3
	// 2 ^ 3 = 1
	// 1 ^ 2 = 3
	// 1 ^ 2 ^ 3 = 0
	k := 2
	expect := 6
	runSample(t, a, k, expect)
}

func TestSampleN(t *testing.T) {
	n, k := 100, 100

	arr := make([]int, n)
	for d := 0; d < 3; d++ {

		for i := 0; i < n; i++ {
			arr[i] = int(rand.Int31())
		}
		rand.Shuffle(n, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		expect := bruteForce(arr, k)
		runSample(t, arr, k, expect)
	}
}
