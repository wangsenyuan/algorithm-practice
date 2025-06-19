package main

import "testing"

func runSample(t *testing.T, x int, d int) {
	res := solve(x, d)

	var sum int
	for i := 0; i < len(res); {
		j := i
		for i < len(res) && res[i]-res[j] < d {
			i++
		}
		sum += 1<<(i-j) - 1
	}
	if sum != x {
		t.Fatalf("Sample expect %d, but got %d", x, sum)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 150000, 1)
}
