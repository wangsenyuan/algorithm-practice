package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, d, res := drive(reader)
	if d != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, d)
	}
	if expect < 0 {
		return
	}
	n := len(a)
	m := len(a[0])
	var arr []pair
	box := []int{n, m, -1, -1}
	for i := range n {
		for j := range m {
			if a[i][j] == 'w' && res[i][j] != 'w' {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if res[i][j] == 'w' || res[i][j] == '+' {
				arr = append(arr, pair{i, j})
				box[0] = min(box[0], i)
				box[1] = min(box[1], j)
				box[2] = max(box[2], i)
				box[3] = max(box[3], j)
			}
		}
	}

	if !check(box[0], box[1], box[2], box[3], arr) {
		t.Fatalf("Sample result %v, not a valid frame", res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 8
..w..w..
........
........
..w..w..
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 6
......
.w....
......
..w...
......
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 4
....
.w..
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 6
w..w.w
...w..
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `9 4
....
....
....
....
....
..w.
....
....
.w..
`
	expect := 4
	runSample(t, s, expect)
}
