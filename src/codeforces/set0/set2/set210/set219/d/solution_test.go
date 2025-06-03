package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, ans := process(reader)
	y, cnt := readTwoNums(reader)
	expect := readNNums(reader, cnt)
	if x != y || !reflect.DeepEqual(ans, expect) {
		t.Fatalf("Sample expect %d %v, but got %d %v", y, expect, x, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 1
2 3
0 1
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
1 4
2 4
3 4
2 3
1 2 3`
	runSample(t, s)
}
