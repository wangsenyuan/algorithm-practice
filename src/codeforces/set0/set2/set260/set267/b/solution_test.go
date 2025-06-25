package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	dominoes, res := process(reader)
	expect := readString(reader)

	if len(res) == 0 || (expect == "No solution" != (res[0] == "No solution")) {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}

	if expect == "No solution" {
		return
	}

	if len(res) != len(dominoes) {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
	n := len(dominoes)
	ans := make([][]int, n)
	for i, x := range res {
		var id int
		pos := readInt([]byte(x), 0, &id) + 1
		id--
		if x[pos] == '-' {
			dominoes[id][0], dominoes[id][1] = dominoes[id][1], dominoes[id][0]
		}
		ans[i] = dominoes[id]
	}

	// t.Logf("Sample got:  %v", ans)

	for i := 1; i < n; i++ {
		if ans[i-1][1] != ans[i][0] {
			t.Fatalf("Sample result given wrong arrangement %v", ans)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2
2 4
2 4
6 4
2 1
2 -
1 -
5 -
3 +
4 -
`)
}


func TestSample2(t *testing.T) {
	runSample(t, `1
0 0
1 +
`)
}
