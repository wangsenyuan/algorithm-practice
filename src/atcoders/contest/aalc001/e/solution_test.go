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
	runSample(t, `5 5
1 2 3 4 5
RGBGB
2 1 5
1 2 10
2 1 5
2 2 4
2 1 3
`, []int{36, 100, 0, 30})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 5
1000000000 1000000000 1000000000
RGB
2 1 3
1 2 998244353
2 1 3
1 2 998244354
2 1 3
`, []int{747068572, 0, 716070898})
}
