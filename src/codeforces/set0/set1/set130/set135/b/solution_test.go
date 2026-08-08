package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	points, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	check := func(a Point, b Point, c Point) bool {
		v1 := a.sub(b)
		v2 := c.sub(b)
		return v1.dot(v2) == 0
	}

	checkRect := func(arr []int) bool {
		for i := 0; i < 4; i++ {
			a := points[(arr[(i+3)%4] - 1)]
			b := points[(arr[i] - 1)]
			c := points[(arr[(i+1)%4] - 1)]
			if !check(a, b, c) {
				return false
			}
		}
		return true
	}

	for _, cur := range res {
		if !checkRect(cur) {
			t.Fatalf("Sample rect not valid %v", cur)
		}
	}
	d1 := points[res[0][0]-1].sub(points[res[0][1]-1])
	d2 := points[res[0][1]-1].sub(points[res[0][2]-1])
	if d1.len2() != d2.len2() {
		t.Fatalf("Sample rect not square %v", res[0])
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0 0
10 11
10 0
0 11
1 1
2 2
2 1
1 2
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `0 0
1 1
2 2
3 3
4 4
5 5
6 6
7 7
	`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `0 0
4 4
4 0
0 4
1 2
2 3
3 2
2 1
`, true)
}

func TestRegressionTwoRightAnglesDoNotMakeRectangle(t *testing.T) {
	runSample(t, `0 0
1 0
0 3
-1 1
10 10
11 10
11 11
10 11
`, false)
}
