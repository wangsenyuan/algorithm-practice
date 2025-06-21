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
	runSample(t, `2
1 2 3 4
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 3 4 6 3 4 100 200
5`)
}

func TestSample3(t *testing.T) {
	runSample(t, `17
814 744 145 886 751 1000 272 914 270 529 467 164 410 369 123 424 991 12 702 582 561 858 746 950 598 393 606 498 648 686 455 873 728 858
318`)
}
