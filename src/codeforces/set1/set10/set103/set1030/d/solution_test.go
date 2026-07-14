package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func area2(pts [][]int) int {
	x1, y1 := pts[0][0], pts[0][1]
	x2, y2 := pts[1][0], pts[1][1]
	x3, y3 := pts[2][0], pts[2][1]
	v := (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	if v < 0 {
		return -v
	}
	return v
}

func runSample(t *testing.T, s string, expectOK bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	res := solve(n, m, k)
	if res.ok != expectOK {
		t.Fatalf("Sample expect ok=%v, but got ok=%v pts=%v", expectOK, res.ok, res.pts)
	}
	if !expectOK {
		return
	}
	if len(res.pts) != 3 {
		t.Fatalf("Sample expect 3 points, but got %v", res.pts)
	}
	for _, p := range res.pts {
		if p[0] < 0 || p[0] > n || p[1] < 0 || p[1] > m {
			t.Fatalf("Point %v out of bounds for n=%d m=%d", p, n, m)
		}
	}
	// Required area is n*m/k, so twice the area is 2*n*m/k.
	if (2*n*m)%k != 0 {
		t.Fatalf("Impossible area for n=%d m=%d k=%d", n, m, k)
	}
	if area2(res.pts) != 2*n*m/k {
		t.Fatalf("Sample expect area %d, but got %d from %v", 2*n*m/k, area2(res.pts), res.pts)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 3 3
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4 7
`, false)
}
