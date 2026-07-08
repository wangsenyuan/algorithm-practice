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
	runSample(t, `3 4
4 3 2
1 1 7
2 3 5
1 2 0
2 4 2
`, []int{11, 12})
}

func TestSample2(t *testing.T) {
	runSample(t, `8 10
320 578 244 604 145 839 156 857
2 400 556
1 5 168
2 254 62
2 145 301
1 1 23
1 3 0
2 413 758
2 297 613
1 8 451
2 598 692
`, []int{3824, 2032, 2073, 4350, 3596, 4884})
}
