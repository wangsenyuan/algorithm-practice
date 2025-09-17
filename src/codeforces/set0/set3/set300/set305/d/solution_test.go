package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 8 2
1 2
2 3
3 4
3 6
4 5
4 7
5 6
6 7
`
	expect := 2
	// 因为 3, 6已经有边了，所以可以在1...4, 2....5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 0 2`
	// 1, 2, 3, 4, 5, 6, 7
	// 1 -> 4 -> 7
	// 2 -> 5
	// 3 -> 6
	expect := 12
	runSample(t, s, expect)
}

// 1， 2， 3， 4， 5， 6， 7， 8 （k = 2)
// 如果在 1...4 和 5 ...8中间都加一条边，是不行的, 这时候 1...8的距离 = 3 了
// 也就是说，只能选择一个区间 (1....k+1)

func TestSample3(t *testing.T) {
	s := `7 1 3
1 5`
	// 1, 2, 3, 4, 5, 6, 7
	expect := 4
	runSample(t, s, expect)
}
