package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
4 1
2 3 1 4
`, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2 1
2 4
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
6 1
3 9 1 6 4 5
`, 40)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
7 1
1 2 3 67 13 8 8
`, 64)
}
