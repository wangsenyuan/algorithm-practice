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
3 1
1 2 3
`, []int64{3})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
5 1
1 4 5 2 6
`, []int64{15})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
6 2
1 1 4 5 1 4
`, []int64{26})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
10 2
230 24 3 42 432 234 934 2389 333 444
`, []int64{8590})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
3 1
100000000 100000000 100000000
`, []int64{0})
}
