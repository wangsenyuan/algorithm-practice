package main

import "testing"

func runSample(t *testing.T, a []string, flag int) {
	n := len(a)
	m := len(a[0])
	var r, c int

	var cnt int

	ask := func(cmd string) []int {
		cnt++
		if cnt > 2*n*m {
			t.Fatalf("asked too much times %d", cnt)
		}
		switch cmd {
		case "L":
			if flag&1 == 1 {
				cmd = reverse(cmd)
			}
		case "R":
			if flag&1 == 1 {
				cmd = reverse(cmd)
			}
		case "U":
			if flag&2 == 2 {
				cmd = reverse(cmd)
			}
		case "D":
			if flag&2 == 2 {
				cmd = reverse(cmd)
			}
		}

		switch cmd {
		case "L":
			c = max(c-1, 0)
		case "R":
			c = min(c+1, m-1)
		case "U":
			r = max(r-1, 0)
		default:
			r = min(r+1, n-1)
		}
		if a[r][c] == '*' {
			t.Fatalf("bad luck, reach a dangerous place at (%d, %d)", r, c)
		}
		return []int{r, c}
	}

	solve(a, ask)

	if a[r][c] != 'F' {
		t.Fatalf("bad luck, not reach the destination at (%d, %d)", r, c)
	}
}

func TestSample1(t *testing.T) {
	a := []string{
		"...",
		"**.",
		"F*.",
		"...",
	}
	runSample(t, a, 3)
}
