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
	runSample(t, `3
1 2
2 3
`, []int{3, 3, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 2
2 3
2 4
1 5
`, []int{4, 5, 5, 5, 4})
}
