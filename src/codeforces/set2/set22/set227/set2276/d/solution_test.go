package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 1
1 2
`, "Brain Melt")
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 2
1 3
`, "2")
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
1 2
2 3
1 3
`, "Many")
}

func TestSample4(t *testing.T) {
	runSample(t, `4 3
1 2
2 3
3 4
`, "Many")
}
