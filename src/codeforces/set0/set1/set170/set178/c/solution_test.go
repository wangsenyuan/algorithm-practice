package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 2 7
+ 11 0
+ 22 2
+ 33 6
+ 44 0
+ 55 0
- 22
+ 66 0
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1 6
+ 123 0
+ 234 1
+ 345 2
- 234
+ 456 0
+ 567 0
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9 3 20
+ 309 1
+ 321 6
- 321
+ 99 3
+ 217 3
+ 95 3
+ 936 1
- 95
+ 354 8
+ 82 6
+ 653 5
+ 730 7
+ 272 8
- 309
+ 211 4
- 217
+ 385 3
- 385
+ 27 3
- 936
`
	// 0 3 6
	// 1 4 7
	// 2 5 8

	expect := 10
	runSample(t, s, expect)
}
