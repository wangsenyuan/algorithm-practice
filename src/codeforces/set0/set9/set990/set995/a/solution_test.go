package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, park, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	n := len(park[0])

	pos := make([][]int, k+1)
	for i := 1; i <= 2; i++ {
		for j := range n {
			if park[i][j] != 0 {
				pos[park[i][j]] = []int{i, j}
			}
		}
	}
	want := make([][]int, k+1)
	for _, i := range []int{0, 3} {
		for j := range n {
			if park[i][j] != 0 {
				want[park[i][j]] = []int{i, j}
				park[i][j] = 0
			}
		}
	}

	for _, cur := range res {
		id, x, y := cur[0], cur[1], cur[2]
		x--
		y--
		u, v := pos[id][0], pos[id][1]
		if abs(x-u)+abs(y-v) != 1 {
			t.Fatalf("Sample result %v, not valid", res)
		}

		if park[x][y] != 0 {
			t.Fatalf("Sample result %v, move %d to (%d %d), but park[%d][%d] = %d", res, id, x, y, x, y, park[x][y])
		}
		park[u][v] = 0
		park[x][y] = id
		pos[id] = []int{x, y}
	}

	for i := 1; i <= k; i++ {
		if !slices.Equal(pos[i], want[i]) {
			t.Fatalf("Sample result %v, pos[%d] = %v, want[%d] = %v", res, i, pos[i], i, want[i])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
1 2 0 4
1 2 0 4
5 0 0 3
0 5 0 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2
1
2
1
2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 2
1
1
2
2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 2
1 0
0 2
0 1
0 2
`
	expect := true
	runSample(t, s, expect)
}