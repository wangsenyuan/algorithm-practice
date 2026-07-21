package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5
0 0 0 0 0
`, []int64{0})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
4
0 13 15 8
`, []int64{59})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
4
13 15 0 8
`, []int64{64})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
8
1 2 3 4 5 6 7 8
`, []int64{72})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
5
99999999 100000000 99999999 99999999 99999999
`, []int64{899999993})
}

func TestSample6(t *testing.T) {
	runSample(t, `1
5
2 3 4 3 2
`, []int64{24})
}
