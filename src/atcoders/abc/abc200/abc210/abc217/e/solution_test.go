package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `8
1 4
1 3
1 2
1 1
3
2
1 0
2
`, []int{1, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `9
1 5
1 5
1 3
2
3
2
1 6
3
2
`, []int{5, 3, 5})
}
