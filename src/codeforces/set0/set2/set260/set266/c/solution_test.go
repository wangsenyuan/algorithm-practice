package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	ops, n, ones := drive(reader)
	if len(ops) > 100000 {
		t.Fatalf("too many operations: %d", len(ops))
	}

	mat := make([][]int, n+1)
	for i := range mat {
		mat[i] = make([]int, n+1)
	}
	for _, p := range ones {
		mat[p[0]][p[1]] = 1
	}

	for _, op := range ops {
		if len(op) != 3 {
			t.Fatalf("bad op %v", op)
		}
		typ, i, j := op[0], op[1], op[2]
		if i < 1 || i > n || j < 1 || j > n || i == j || typ < 1 || typ > 2 {
			t.Fatalf("invalid op %v", op)
		}
		if typ == 1 {
			mat[i], mat[j] = mat[j], mat[i]
		} else {
			for r := 1; r <= n; r++ {
				mat[r][i], mat[r][j] = mat[r][j], mat[r][i]
			}
		}
	}

	var cnt int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			cnt++
			if i <= j {
				t.Fatalf("one at (%d,%d) is not strictly below diagonal after ops %v", i, j, ops)
			}
		}
	}
	if cnt != n-1 {
		t.Fatalf("expected %d ones, found %d", n-1, cnt)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
1 2
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
3 1
1 3
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
2 1
3 2
`)
}
