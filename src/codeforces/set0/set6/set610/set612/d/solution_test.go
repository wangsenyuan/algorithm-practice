package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	var expect int
	fmt.Fscan(reader, &expect)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	for _, cur := range res {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		if cur[0] != l || cur[1] != r {
			t.Fatalf("Sample expect %d %d, but got %v", l, r, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
0 5
-3 2
3 8
2
0 2
3 5
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2
0 5
-3 3
3 8
1
0 5
`
	runSample(t, s)
}
