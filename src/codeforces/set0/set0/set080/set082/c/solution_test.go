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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
40 10 30 20
1 2 1
2 3 1
4 2 1
`
	expect := "0 1 3 2"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
5 4 3 2 1
1 2 1
2 3 1
2 4 1
4 5 1
`
	expect := "0 1 4 2 3"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
3 1 2
1 3 1
3 2 1
`
	expect := "0 2 1"
	runSample(t, s, expect)
}
