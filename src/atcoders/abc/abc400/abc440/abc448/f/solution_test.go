package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	points, res := drive(reader)

	n := len(points)
	if len(res) != n {
		t.Fatalf("Sample expect a %d length path, but got %v", n, res)
	}
	if res[0] != 1 {
		t.Fatalf("Sample expect start and end at 1, but got %v", res)
	}

	var sum int

	vis := make([]bool, n+1)

	for i := 1; i < n; i++ {
		if vis[res[i]] {
			t.Fatalf("Sample result %v, not valid, point %d visited twice", res, res[i])
		}
		vis[res[i]] = true
		tmp := distance(points[res[i-1]-1], points[res[i]-1])
		sum += tmp
	}

	sum += distance(points[res[n-1]-1], points[0])

	if sum > 1e10 {
		t.Fatalf("Sample result %v, not valid, sum %d is too large", res, sum)
	}
}

func distance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int {
	return max(x, -x)
}

func TestSample1(t *testing.T) {
	s := `10
9706344 19786176
19341349 15565412
5711023 19068083
12521132 14054301
14767612 17088029
14961700 18526945
13801766 5740101
6581153 8643675
13176196 16586661
4086263 5172719
`
	runSample(t, s)
}
