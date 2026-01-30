package main

import "testing"

func runSample(t *testing.T, s []string) {
	var cnt int
	n := len(s)

	que := make([]int, n*n)
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, n)
	}

	check := func(r1, c1, r2, c2 int) bool {
		if s[r1][c1] == '#' {
			return false
		}
		var head, tail int
		que[head] = r1*n + c1
		head++
		for tail < head {
			r, c := que[tail]/n, que[tail]%n
			tail++
			if r+1 <= r2 && !marked[r+1][c] && s[r+1][c] == '.' {
				marked[r+1][c] = true
				que[head] = (r+1)*n + c
				head++
			}
			if c+1 <= c2 && !marked[r][c+1] && s[r][c+1] == '.' {
				marked[r][c+1] = true
				que[head] = r*n + c + 1
				head++
			}
		}
		ok := marked[r2][c2]
		for i := range head {
			r, c := que[i]/n, que[i]%n
			marked[r][c] = false
		}
		return ok
	}

	ask := func(res []int) bool {
		r1, c1, r2, c2 := res[0], res[1], res[2], res[3]
		if r1 > r2 || c1 > c2 {
			t.Fatalf("Sample ask wrong %v", res)
		}
		cnt++
		if cnt > 4*n {
			t.Fatalf("Sample ask too much %d", cnt)
		}
		if r2-r1+c2-c1 < n-1 {
			t.Fatalf("Sample ask wrong %v too short", res)
		}
		return check(r1-1, c1-1, r2-1, c2-1)
	}

	res := solve(n, ask)
	if len(res) != n+n-2 {
		t.Fatalf("Sample result %v, not valid", res)
	}
	var r, c int
	for _, ch := range res {
		if ch == 'D' {
			r++
		} else {
			c++
		}
		if r >= n || c >= n || s[r][c] == '#' {
			t.Fatalf("Sample result %v, not valid", res)
		}
	}
	if r != n-1 || c != n-1 {
		t.Fatalf("Sample result %v, not valid", res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"..#.",
		"#...",
		"###.",
		"....",
	}
	runSample(t, s)
}
