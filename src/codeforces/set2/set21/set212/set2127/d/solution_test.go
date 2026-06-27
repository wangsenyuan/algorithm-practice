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
2 1
1 2
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
3 3
1 2
1 3
2 3
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
5 4
1 2
1 3
3 4
3 5
`, 8)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
4 3
1 2
1 3
1 4
`, 12)
}
