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
	runSample(t, `4
2 4 6 10
5
2 3
1 2
3 2
2 4
3 1
`, []int{4, 1})
}

func TestHalfIntegerAtRightEndpointIsOutside(t *testing.T) {
	runSample(t, `1
2
1
3 2
`, []int{0})
}
