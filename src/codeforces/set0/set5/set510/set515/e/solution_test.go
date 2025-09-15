package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
2 2 2 2 2
3 5 2 1 4
1 3
2 2
4 5
`
	expect := []int{12, 16, 18}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
5 1 4
5 1 4
3 3
2 2
1 1
`
	expect := []int{17, 22, 11}
	runSample(t, s, expect)
}
