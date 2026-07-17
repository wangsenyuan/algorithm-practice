package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
1 2
2 3
3 4
2 5
5 6
`, "Yes")
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 2
2 3
3 4
2 5
3 6
`, "No")
}
