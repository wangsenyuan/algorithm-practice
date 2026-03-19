package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	rects, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	var sum int
	box := []int{inf, inf, -inf, -inf}

	for _, i := range res {
		i--
		x1, y1, x2, y2 := rects[i][0], rects[i][1], rects[i][2], rects[i][3]
		sum += (x2 - x1) * (y2 - y1)
		box[0] = min(box[0], x1)
		box[1] = min(box[1], y1)
		box[2] = max(box[2], x2)
		box[3] = max(box[3], y2)
	}

	w := box[2] - box[0]
	h := box[3] - box[1]
	if sum != w*h || w != h {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

const inf = 1 << 60

func TestSample1(t *testing.T) {
	s := `9
0 0 1 9
1 0 9 1
1 8 9 9
8 1 9 8
2 2 3 6
3 2 7 3
2 6 7 7
5 3 7 6
3 3 5 6
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
0 0 1 9
1 0 9 1
1 8 9 9
8 1 9 8
`
	expect := false
	runSample(t, s, expect)
}
