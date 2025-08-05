package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `aaab
2
aa 0 0
aab 1 1
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `ltntlnen
3
n 0 0
ttlneenl 1 4
lelllt 1 1
`, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, `a
0
`, 1)
}
