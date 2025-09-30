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
	for i := range expect {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		if res[i][0] != n || res[i][1] != m {
			t.Fatalf("Sample result %v is not correct, expect (%d %d)", res, n, m)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `26
6
1 26
2 9
3 5
5 3
9 2
26 1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
2
1 2
2 1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `8
4
1 8
2 3
3 2
8 1`
	runSample(t, s)
}
