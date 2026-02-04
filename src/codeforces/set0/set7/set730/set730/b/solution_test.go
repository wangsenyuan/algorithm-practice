package main

import (
	"math/rand"
	"slices"
	"testing"
)

func runSample(t *testing.T, a []int) {
	n := len(a)
	var cnt int
	ask := func(i int, j int) string {
		cnt++
		if cnt > (3*n+1)/2-2 {
			t.Fatalf("Sample asked too much times %d", cnt)
		}
		i--
		j--
		if a[i] < a[j] {
			return "<"
		}
		if a[i] > a[j] {
			return ">"
		}
		return "="
	}
	res := solve(n, ask)
	i := res[0] - 1
	j := res[1] - 1
	if a[i] != slices.Min(a) {
		t.Fatalf("Sample result %v, not correct", res)
	}
	if a[j] != slices.Max(a) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	runSample(t, a)
}

func TestSample3(t *testing.T) {
	a := make([]int, 50)
	for i := range a {
		a[i] = i + 1
	}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	runSample(t, a)
}

func TestSample4(t *testing.T) {
	a := make([]int, 50)
	rng := rand.New(rand.NewSource(1))
	for i := range a {
		a[i] = rng.Intn(1000) + 1
	}

	rng.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	runSample(t, a)
}

func TestSample5(t *testing.T) {
	a := []int{1, 1, 1}
	runSample(t, a)
}

func TestSample6(t *testing.T) {
	a := []int{3, 2, 1}
	runSample(t, a)
}
