package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	sz, res := process(reader)
	expect_sz := readNum(reader)
	expect_res := readNNums(reader, 2)

	if sz != expect_sz || !reflect.DeepEqual(res, expect_res) {
		t.Fatalf("Sample expect %d, %v, but got %d, %v", expect_sz, expect_res, sz, expect_res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 4
ABDC
ABDC
3
2 1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 6
ABCCBA
ABCCBA
1
2 6`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `16 4
BAAC
BACA
ACBC
ABCC
CCAC
BBCC
CCAB
ABCC
CBCA
BCBC
BCBC
CBBB
BBAA
BACA
ABCB
AABA
9
2 4`
	runSample(t, s)
}
