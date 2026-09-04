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
	if res := drive(reader); !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
1
2
3
4
5
`, []int{1, 2, 3, 5, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `7 7
7
7
7
7
7
7
7
`, []int{1, 2, 3, 4, 5, 7, 6})
}

func TestSample3(t *testing.T) {
	runSample(t, `10 6
1
5
2
9
6
6
`, []int{1, 2, 3, 4, 5, 7, 6, 8, 10, 9})
}
