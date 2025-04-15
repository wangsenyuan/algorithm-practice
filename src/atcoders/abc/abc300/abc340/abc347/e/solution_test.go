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
	s := `3 4
1 3 3 2
`
	expect := "6 2 2"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 6
1 2 3 2 4 2
`
	expect := "15 9 12 7"
	runSample(t, s, expect)
}