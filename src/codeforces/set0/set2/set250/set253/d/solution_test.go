package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	ans := readNum(reader)
	if res != ans {
		t.Errorf("Sample expect %d, but got %d", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 4
aabb
baab
baab
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 5 1
ababa
ccaca
ccacb
cbabc
1`
	runSample(t, s)
}
