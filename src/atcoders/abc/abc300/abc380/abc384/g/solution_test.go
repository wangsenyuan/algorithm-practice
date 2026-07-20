package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int64) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
2 4
3 5
4
1 1
1 2
2 1
2 2
`, []int64{1, 4, 2, 6})
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1163686 28892 1263085 2347878 520306
1332157 1202905 2437161 1291976 563395
5
5 3
1 5
2 3
1 2
5 5
`, []int64{13331322, 2209746, 6366712, 207690, 20241215})
}
