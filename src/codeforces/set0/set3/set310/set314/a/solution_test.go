package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if !reflect.DeepEqual(ans, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5 0
5 3 4 1 2
`
	expect := []int{2, 3, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 -10
5 5 1 7 5 1 2 4 9 2
`
	expect := []int{2,
		4,
		5,
		7,
		8,
		9,
	}
	runSample(t, s, expect)
}
