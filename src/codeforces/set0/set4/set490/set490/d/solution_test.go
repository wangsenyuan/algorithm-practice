package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, sz1, sz2 := process(reader)

	expect := readNum(reader)

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}

	if cnt < 0 {
		return
	}

	if sz1[0]*sz1[1] != sz2[0]*sz2[1] {
		t.Fatalf("Sample result %v, %v not correct", sz1, sz2)
	}
}

func TestSample1(t *testing.T) {
	s := `2 6
2 3
1`
	runSample(t, s)
}
func TestSample2(t *testing.T) {
	s := `36 5
10 16
3`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 5
2 1
-1`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `3 6
2 1
4`
	runSample(t, s)
}
