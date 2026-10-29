package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 8
3 1 3 2
3
1 0
3
2 3
3
2 3
3
1 2
`, []int{26, 35, 20, 8})
}

func TestSample2(t *testing.T) {
	runSample(t, `1 5
998244353
3
2 998244353
3
1 1000000000
3
`, []int{0, 0, 1755647})
}

func TestSample3(t *testing.T) {
	runSample(t, `3 2
1 1 3
1 1
3
`, []int{18})
}

func TestSample4(t *testing.T) {
	runSample(t, `3 2
1 1 3
2 1
3
`, []int{7})
}
