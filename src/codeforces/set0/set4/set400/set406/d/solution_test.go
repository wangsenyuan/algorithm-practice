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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, `6
1 4
2 1
3 2
4 3
6 4
7 4
3
3 1
5 6
2 3
5 6 3
`)
}
