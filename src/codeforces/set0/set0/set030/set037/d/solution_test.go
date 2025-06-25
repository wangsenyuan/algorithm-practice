package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 1 1
1 1 1
6`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 1 1
1 2 3
36`)

	// 对于用户1， (1, 1), (1, 2), (1, 3)
	// 对于用户2,  (2, 2), (2, 3)
	// 对于用户3  （3， 3）
}
