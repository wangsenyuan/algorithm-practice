package main

import (
	"math/rand"
	"slices"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	n := len(arr) / 2
	freq := make([]int, 2*n+1)

	var cnt int
	ask := func(req []int) int {
		cnt++
		if cnt > 3*n {
			t.Fatalf("Sample solution ask too much times %d", cnt)
		}
		var todo []int
		for _, i := range req {
			freq[arr[i-1]]++
			todo = append(todo, arr[i-1])
		}
		var ans int

		for _, v := range todo {
			if freq[v] == 2 && v > ans {
				ans = v
			}
		}
		for _, v := range todo {
			freq[v] = 0
		}
		return ans
	}

	res := solve(n, ask)

	if !slices.Equal(res, arr) {
		t.Fatalf("Sample not solved, got %v", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 2, 1, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 1, 2})
}

func TestSample3(t *testing.T) {
	n := 100
	arr := make([]int, n*2)
	for i := range n {
		arr[i] = i + 1
		arr[i+100] = i + 1
	}

	rand.Shuffle(2*n, func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	runSample(t, arr)
}
