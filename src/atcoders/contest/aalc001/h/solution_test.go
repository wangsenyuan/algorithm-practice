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
	runSample(t, `4 5
8 0 11 5
2 1 4 998244350
1 2 6
2 2 4 10
1 4 0
2 3 4 1000000000
`, []int{21, 32, 1755658})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 7
650000000 0 910000007 123456789 0 870000013
2 3 4 750000021
2 2 6 987654321
1 5 420000009
2 3 6 314159265
1 1 0
2 1 3 271828182
2 6 6 1000000000
`, []int{1783456817, 894622424, 2637616083, 1181828189, 1870000013})
}
