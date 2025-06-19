package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := process(reader)
	expect := readString(reader)

	if len(res) > 0 != (expect == "YES") {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}

	if expect != "YES" {
		return
	}
	_, y := res[0], res[1]

	get := func(v int) (r int, c int) {
		r, c = v/y, v%y
		return
	}

	for i := 1; i < len(a); i++ {
		r1, c1 := get(a[i-1] - 1)
		r2, c2 := get(a[i] - 1)
		if abs(r2-r1)+abs(c2-c1) != 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `8
1 2 3 6 9 8 5 2
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
1 2 1 2 5 3
NO`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2
1 10
YES`)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
1 2 2
NO`)
}
