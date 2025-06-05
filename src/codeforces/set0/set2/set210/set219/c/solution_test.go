package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans, res, s := process(reader)
	expect := readNum(reader)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
	var cnt int
	for i := range len(s) {
		if s[i] != res[i] {
			cnt++
		}
		if i > 0 && res[i] == res[i-1] {
			t.Fatalf("Sample result %s, not correct", res)
		}
	}
	if cnt != ans {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 3
ABBACC
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
BBB
1`)
}
