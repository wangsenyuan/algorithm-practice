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
	runSample(t, `9 8 4
4216793+
717*4*54
7+5+727*
45149+71
8+26697*
+189*2+9
5+*244+7
42595952
97+*315+
67
420
3
727
`, []int{2, 0, 4, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `2 22 2
1+64*15625************
111111+222222+666666+1
999999
1000000
`, []int{2, 3})
}
