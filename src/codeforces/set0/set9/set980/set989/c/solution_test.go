package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int) {
	res := solve(nums[0], nums[1], nums[2], nums[3])

	freq := make([]int, 4)
	n := len(res)
	m := len(res[0])
	if n > 50 || m > 50 {
		t.Fatalf("Sample result took too much space %d %d", n, m)
	}
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}

	var visit func(r int, c int)
	visit = func(r int, c int) {
		marked[r][c] = true
		if r > 0 && res[r-1][c] == res[r][c] && !marked[r-1][c] {
			visit(r-1, c)
		}
		if c > 0 && res[r][c-1] == res[r][c] && !marked[r][c-1] {
			visit(r, c-1)
		}
		if r < n-1 && res[r+1][c] == res[r][c] && !marked[r+1][c] {
			visit(r+1, c)
		}
		if c < m-1 && res[r][c+1] == res[r][c] && !marked[r][c+1] {
			visit(r, c+1)
		}
	}

	for i := range n {
		for j := range m {
			if !marked[i][j] {
				visit(i, j)
				freq[int(res[i][j]-'A')]++
			}
		}
	}

	if !reflect.DeepEqual(freq, nums) {
		t.Errorf("Sample expect %v, but got %v", nums, freq)
	}

}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 1, 1, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{5, 3, 2, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1, 6, 4, 5})
}

func TestSample4(t *testing.T) {
	runSample(t, []int{100, 100, 100, 100})
}
