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
5
3 3 2 4 3
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
7
9 5 7 7 4 7 7
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
9
1 1 1 1 1 1 1 1 1
`, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
3
2 2 2
`, 3)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
5
1 2 3 4 5
`, 1)
}

func TestSample6(t *testing.T) {
	runSample(t, `1
5
2 1 3 2 2
`, 3)
}

func TestSample7(t *testing.T) {
	runSample(t, `1
7
2 2 1 2 3 2 2
`, 5)
}

func TestSample8(t *testing.T) {
	runSample(t, `1
9
2 1 2 3 2 1 2 3 2
`, 5)
}
