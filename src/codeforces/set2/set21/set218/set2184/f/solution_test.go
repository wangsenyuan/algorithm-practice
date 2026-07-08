package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	for i := range res {
		if !strings.EqualFold(res[i], expect[i]) {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
4
1 2
1 3
1 4
3
1 2
1 3
9
1 2
3 1
2 4
5 2
5 6
3 7
8 3
8 9
`, []string{"YES", "NO", "YES"})
}
