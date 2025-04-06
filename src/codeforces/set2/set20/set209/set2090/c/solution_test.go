package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	reader = bufio.NewReader(strings.NewReader(expect))
	for _, cur := range res {
		tmp := readNNums(reader, 2)
		if tmp[0] != cur[0] || tmp[1] != cur[1] {
			t.Errorf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6
0 1 1 0 0 1`
	expect := `1 1
1 2
2 1
1 4
4 1
1 5`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 0 0 1 1`
	expect := `1 1
1 4
4 1
1 2
2 1`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
0 1 1 0 0 1 0`
	expect := `1 1
1 2
2 1
1 4
4 1
1 5
1 7`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
0 1 1 0 0 1 0 1 0 0`
	expect := `1 1
1 2
2 1
1 4
4 1
1 5
1 7
2 2
4 4
7 1`
	runSample(t, s, expect)
}
