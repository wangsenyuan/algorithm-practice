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

	if expect == "-1" {
		if len(res) > 0 {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
	reader = bufio.NewReader(strings.NewReader(expect))
	ans := readNNums(reader, len(res))

	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
-2 1
`
	expect := "3 1 2 "
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 1 1 1
`
	expect := "1 2 3 4 5"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
-1 2 2
`
	expect := "-1"
	runSample(t, s, expect)
}
