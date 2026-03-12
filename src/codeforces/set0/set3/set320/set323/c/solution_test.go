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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
3 1 2
3 2 1
1
1 2 3 3
`
	expect := []int{1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
4 3 2 1
2 3 4 1
3
1 2 3 4
1 3 2 1
1 4 2 3
`
	expect := []int{1, 1, 2}
	runSample(t, s, expect)
}
