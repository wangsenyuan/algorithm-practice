package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNNums(reader, len(res))

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
3 1 4 1 5 9
3
7 7
0 2
8 17
5 10 18
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
1 500000000000000000
2
1000000000000000000 1000000000000000000
0 1000000000000000000
2 1500000000000000000
`)
}
