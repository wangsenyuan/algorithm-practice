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
	runSample(t, `5
2 4 1 4
3 3 3 5
1 3 4 6
4 5 3 5
5 5 4 6
`, []int{
		3999983,
		3999976,
		3999982,
		3999978,
		3999977,
	})
}
