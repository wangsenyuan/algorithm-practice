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
	s := `4 3 2
1 1 1 1
1 1 1
`
	expect := "0 1 1 0"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 5
1 2 3
4
`
	expect := "0 1 2"
	runSample(t, s, expect)
}