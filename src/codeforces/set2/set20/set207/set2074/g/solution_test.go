package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSamples(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	if tc != len(expect) {
		t.Fatalf("input has %d cases, expect slice has %d", tc, len(expect))
	}
	for i := range tc {
		res := drive(reader)
		if res != expect[i] {
			t.Fatalf("case %d expect %d, but got %d", i+1, expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSamples(t, `6
3
1 2 3
4
2 1 3 4
6
2 1 2 1 1 1
6
1 2 1 3 1 5
9
9 9 8 2 4 4 3 5 3
9
9 9 3 2 4 4 8 5 3
`, []int{6, 24, 5, 30, 732, 696})
}
