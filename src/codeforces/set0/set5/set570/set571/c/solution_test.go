package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	clauses, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for _, cur := range clauses {
		var val int
		for _, v := range cur {
			u := abs(v) - 1
			x := int(res[u] - '0')
			if v > 0 {
				val = x
			} else {
				val = x ^ 1
			}
			if val == 1 {
				break
			}
		}
		if val == 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	s := `2 2
2 1 -2
2 2 -1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
1 1
1 2
3 -1 -2 3
1 -3
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 6
2 1 2
3 1 -2 3
4 -3 5 4 6
2 -6 -4
1 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 5
3 2 3 -5
1 4
2 -1 -4
2 1 -2
2 -3 5
`
	expect := true
	runSample(t, s, expect)
}
