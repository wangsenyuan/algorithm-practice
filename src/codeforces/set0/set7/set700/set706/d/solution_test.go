package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10
+ 8
+ 9
+ 11
+ 6
+ 1
? 3
- 8
? 3
? 8
? 11
`
	expect := []int{11, 10, 14, 13}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
? 1
+ 1
+ 8
- 1
+ 2
+ 7
+ 4
+ 7
+ 3
? 7
`
	expect := []int{1, 15}
	runSample(t, s, expect)
}
