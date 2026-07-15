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
	runSample(t, `5
2 3 5 4 13
`, []int{2, 3, 5, 7, 11})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
10 3 7
`, []int{10, 3, 7})
}

func TestSample3(t *testing.T) {
	runSample(t, `3
8168 59292 95740
`, []int{8168, 59293, 3})
}
