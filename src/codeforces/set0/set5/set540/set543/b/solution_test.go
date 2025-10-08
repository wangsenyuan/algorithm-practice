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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 4
1 2
2 3
3 4
4 5
1 3 2
3 5 2`
	runSample(t, s, 0)
}

func TestSample2(t *testing.T) {
	s := `5 4
1 2
2 3
3 4
4 5
1 3 2
2 4 2`
	runSample(t, s, 1)
}

func TestSample3(t *testing.T) {
	s := `5 4
1 2
2 3
3 4
4 5
1 3 2
3 5 1`
	runSample(t, s, -1)
}

func TestSample4(t *testing.T) {
	s := `9 9
1 2
2 3
2 4
4 5
5 7
5 6
3 8
8 9
9 6
1 7 4
3 6 3`
	runSample(t, s, 2)
}

func TestSample5(t *testing.T) {
	s := `10 11
1 3
2 3
3 4
4 5
4 6
3 7
3 8
4 9
4 10
7 9
8 10
1 5 3
6 2 3`
	runSample(t, s, 6)
}
