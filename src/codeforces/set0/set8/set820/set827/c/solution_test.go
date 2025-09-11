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
		t.Errorf("Sample test failed. Expect: %v, Got: %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `ATGCATGC
4
2 1 8 ATGC
2 2 6 TTT
1 4 T
2 2 6 TA
`
	expect := []int{8, 2, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `GAGTTGTTAA
6
2 3 4 TATGGTG
1 1 T
1 6 G
2 5 9 AGTAATA
1 10 G
2 2 6 TTGT
`
	expect := []int{0, 3, 1}
	runSample(t, s, expect)
}
