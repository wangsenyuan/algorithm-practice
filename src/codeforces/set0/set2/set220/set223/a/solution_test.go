package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	cnt, res := solve(s)
	if cnt != expect {
		t.Errorf("Sample expect %d, but got %d", expect, cnt)
	}
	if expect == 0 {
		return
	}
	if !strings.Contains(s, res) {
		t.Errorf("Sample expect %s in %s", res, s)
	}
	stack := make([]int, len(res))
	var top int
	for i, x := range []byte(res) {
		if x == '[' || x == '(' {
			stack[top] = i
			top++
			if x == '[' {
				cnt--
			}
		} else {
			if top == 0 || !check(res[stack[top-1]], x) {
				t.Fatalf("Sample result %s, not correct", res)
			}
			top--
		}
	}
	if cnt != 0 || top != 0 {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "([])", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "(((", 0)
}
