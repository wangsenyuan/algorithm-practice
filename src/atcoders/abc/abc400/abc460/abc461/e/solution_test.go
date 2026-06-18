package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 1
1 3
2 2
1 1
`
	expect := []int{3, 6, 4, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `300000 1
2 300000
`
	expect := []int{0}
	runSample(t, s, expect)
}
