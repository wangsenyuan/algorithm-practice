package main

import (
	"bufio"
	"slices"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, travelers, k, res := drive(reader)
	if k != expect || len(res) != a {
		t.Errorf("Sample expect %d, but got %d, %v", expect, k, res)
	}

	var pos []int

	sort.Ints(res)
	var j int

	for i, cur := range travelers {
		if j == len(res) {
			break
		}
		if res[j] == i+1 {
			j++
			pos = append(pos, cur[0], cur[0]+cur[1])
		}
	}

	slices.Sort(pos)
	pos = slices.Compact(pos)

	m := len(pos)
	wait_at := make([][]int, m)
	leave_at := make([][]int, m)
	j = 0
	for i, cur := range travelers {
		if j == len(res) {
			break
		}
		if res[j] == i+1 {
			j++
			l := sort.SearchInts(pos, cur[0])
			r := sort.SearchInts(pos, cur[0]+cur[1])
			leave_at[r] = append(leave_at[r], i)
			wait_at[l] = append(wait_at[l], i)
		}
	}

	var room int
	for i := range m {
		for _, j := range leave_at[i] {
			if room == 0 {
				t.Fatalf("Sample result is wrong, a person %d is leaving an empty bus", j)
			}
			room--
		}
		for _, j := range wait_at[i] {
			room++
			if room > k {
				t.Fatalf("Sample result is wrong, a person %d is entering the bus when the bus is full", j)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
8 9
3 5
1 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
20 40
10 10
15 5
5 15
20 30
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 8
6 5
1 5
5 5
6 1
1 4
4 4
2 2
6 6
`
	expect := 5
	runSample(t, s, expect)
}
