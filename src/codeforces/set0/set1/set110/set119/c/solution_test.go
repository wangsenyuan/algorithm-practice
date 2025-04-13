package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	n, _, k, tasks, res := process(bufio.NewReader(strings.NewReader(s)))

	get := func(arr [][]int) int {
		if len(arr) != n {
			t.Fatalf("expect %d rows, but got %d", n, len(arr))
		}
		var sum int
		for i, cur := range arr {
			tmp, v := tasks[cur[0]-1], cur[1]
			if v < tmp[0] || v > tmp[1] {
				t.Fatalf("Sample result %v, not correct at %d-th row", arr, i)
			}
			sum += v
			if i > 0 {
				prev, pv := tasks[arr[i-1][0]-1], arr[i-1][1]
				if prev[2] >= tmp[2] {
					t.Fatalf("Sample result %v, not correct at %d-th row", arr, i)
				}
				if pv+k != v && pv*k != v {
					t.Fatalf("Sample result %v, not correct at %d-th row", arr, i)
				}
			}
		}
		return sum
	}

	if len(expect) > 0 {
		if len(res) == 0 {
			t.Fatalf("Sample expect %v, but got nothing", expect)
		}

		x := get(expect)
		y := get(res)

		if x != y {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}

		return
	}
	if len(res) > 0 {
		t.Fatalf("Sample expect no solution, but got %v", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 5 2
1 10 1
1 10 2
1 10 3
1 20 4
1 100 5
`
	expect := [][]int{
		{2, 8},
		{3, 10},
		{4, 20},
		{5, 40},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 3
1 3 1
2 4 4
2 3 3
2 2 2
`

	runSample(t, s, nil)
}

