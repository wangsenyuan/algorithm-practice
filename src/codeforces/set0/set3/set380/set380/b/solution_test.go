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
1 4 4 7 1
1 3 1 2 2
2 1 1
2 4 1
2 3 3
`, []int{2, 0, 1})
}

func TestSameValueInsertedTwiceCountsOnce(t *testing.T) {
	runSample(t, `2 3
1 2 1 1 5
1 2 2 2 5
2 1 1
`, []int{1})
}
