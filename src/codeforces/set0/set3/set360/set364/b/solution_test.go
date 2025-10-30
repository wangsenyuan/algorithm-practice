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
	s := `3 2
1 3 10
`
	expect := []int{4, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 5
1 2 3
`
	expect := []int{6, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 10000
10000 9999 1 10000 10000 10000 1 2 3 4
`
	expect := []int{50010, 6}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `50 1
1 2 4 8 16 32 64 128 256 512 1024 2048 4096 8192 9 17 33 65 129 257 513 1025 2049 4097 8193 1 2 4 8 16 32 64 128 256 512 1024 2048 4096 8192 9 17 33 65 129 257 513 1025 2049 4097 8193
`
	expect := []int{65540, 65540}
	runSample(t, s, expect)
}
