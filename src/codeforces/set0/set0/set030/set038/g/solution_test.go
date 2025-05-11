package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	ans := readNNums(reader, len(res))

	if !reflect.DeepEqual(res, ans) {
		t.Errorf("Sample expect %v, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 0
2 1
`
	expect := "2 1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 3
2 3
3 3
`
	expect := "3 2 1"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 3
1 4
4 3
3 1
5 2
`
	expect := "3 1 5 4 2"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
5 4
3 5
2 0
4 2
1 3
`
	expect := "1 4 2 3 5"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `15
10 14
9 10
4 6
5 15
7 13
11 15
15 15
8 13
13 0
6 14
12 11
1 6
3 13
2 4
14 15
`
	expect := "7 15 6 1 2 8 5 4 3 9 11 10 13 14 12"
	runSample(t, s, expect)
}
