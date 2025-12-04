package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	d, a, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	if res[0] == res[2] && res[1] == res[3] {
		t.Fatalf("Sample result %v, not correct", res)
	}

	n := len(a)
	m := len(a[0])

	marked := make([][]bool, n)
	dist := make([][]int, n)
	for i := range n {
		marked[i] = make([]bool, m)
		dist[i] = make([]int, m)
		for j := range m {
			marked[i][j] = false
			dist[i][j] = -1
		}
	}

	que := make([]int, n*m)

	bfs := func(x int, y int) {
		for i := range n {
			for j := range m {
				dist[i][j] = -1
			}
		}

		var head, tail int
		que[head] = x*m + y
		head++
		dist[x][y] = 0
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			if a[r][c] == 'R' {
				marked[r][c] = true
			}
			if dist[r][c] == d {
				continue
			}
			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m && dist[nr][nc] == -1 && a[nr][nc] != 'X' {
					dist[nr][nc] = dist[r][c] + 1
					que[head] = nr*m + nc
					head++
				}
			}
		}
	}

	bfs(res[0]-1, res[1]-1)
	bfs(res[2]-1, res[3]-1)

	for i := range n {
		for j := range m {
			if a[i][j] == 'R' && !marked[i][j] {
				t.Fatalf("Sample result %v, not correct, can't reach rat (%d, %d)", res, i, j)
			}
		}
	}

}

func TestSample1(t *testing.T) {
	s := `4 4 1
XXXX
XR.X
X.RX
XXXX
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 14 5
XXXXXXXXXXXXXX
X....R...R...X
X..R.........X
X....RXR..R..X
X..R...X.....X
XR.R...X.....X
X....XXR.....X
X....R..R.R..X
XXXXXXXXXXXXXX
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 7 1
XXXXXXX
X.R.R.X
X.....X
X..X..X
X..R..X
X....RX
XXXXXXX
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 9 5
XXXXXXXXX
XRRRRRRRX
XRRRRRRRX
XRRRRRRRX
XRRRRRRRX
XRRRRRRRX
XRRRRRRRX
XRRRRRRRX
XRRRRRRRX
XXXXXXXXX
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `9 13 1
XXXXXXXXXXXXX
XX..X.X..XX.X
XX..X....X.XX
X..XRX.X.XXXX
X...R..X....X
X...X..X....X
XX...RX.....X
X....RX.X..XX
XXXXXXXXXXXXX
`
	expect := true
	runSample(t, s, expect)
}
