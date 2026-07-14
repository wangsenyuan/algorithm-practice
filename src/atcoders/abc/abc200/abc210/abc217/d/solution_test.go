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
	runSample(t, `5 3
2 2
1 3
2 2
`, []int{5, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3
1 2
1 4
2 3
`, []int{2})
}

func TestSample3(t *testing.T) {
	runSample(t, `100 10
1 31
2 41
1 59
2 26
1 53
2 58
1 97
2 93
1 23
2 84
`, []int{69, 31, 6, 38, 38})
}
