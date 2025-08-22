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
	s := `3
2
1
1
`
	runSample(t, s, 5)
}

func TestSample2(t *testing.T) {
	s := `5
2
2
2
2
2
`
	runSample(t, s, 10)
}

func TestSample3(t *testing.T) {
	s := `3
1
250000
1
`
	runSample(t, s, 5)
}
func TestSample4(t *testing.T) {
	s := `2
250000
250000
`
	runSample(t, s, 2)
}

func TestSample5(t *testing.T) {
	s := `5
5
4
3
2
1
`
	runSample(t, s, 9)
}

func TestSample6(t *testing.T) {
	s := `3
123
456
789
`
	runSample(t, s, 3)
}

func TestSample7(t *testing.T) {
	s := `20
10
2
7
9
8
7
2
5
6
6
1
6
4
2
5
5
5
3
2
8
`
	runSample(t, s, 71)
}
