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
	if expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
1 2
4
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `8
1 2
2 3
3 8
2 4
4 5
5 6
6 7
24
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
1 2
2 3
3 4
4 5
5 6
4 7
7 8
4 9
9 10
0
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `7
1 4
4 2
3 2
3 5
2 6
6 7
16
	`)
}

func TestSample5(t *testing.T) {
	runSample(t, `7
1 2
2 3
3 4
3 5
4 6
6 7
48
	`)
}

func TestSample6(t *testing.T) {
	runSample(t, `5
3 4
1 2
1 3
2 5
4
	`)
}
