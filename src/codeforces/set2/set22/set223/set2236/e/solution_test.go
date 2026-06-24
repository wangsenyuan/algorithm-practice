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
1
1
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2
1 2
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
3
2 1 1
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
4
2 1 4 3
`, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
5
1 2 4 5 3
`, 1)
}

func TestSample6(t *testing.T) {
	runSample(t, `1
6
3 2 1 6 5 4
`, 3)
}

func TestSample7(t *testing.T) {
	runSample(t, `1
10
1 1 2 3 4 1 6 5 7 8
`, 4)
}
