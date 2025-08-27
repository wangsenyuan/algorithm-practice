package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
6
1 2 1
3 4 1
2 3 2
1 3 2
2 4 1
1 4 2
`, []int{2, 0, 2, 3, 1, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
8
2 2 1
3 3 2
1 1 1
1 3 2
2 3 2
3 3 1
1 2 1
2 2 2
`, []int{2, 2, 1, 3, 3, 2, 0, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `7
10
5 7 1
5 6 2
7 7 2
6 7 2
5 5 1
3 6 2
1 3 2
5 6 1
1 3 1
6 7 1
`, []int{4, 6, 7, 7, 6, 7, 7, 5, 2, 1})
}

func TestSample4(t *testing.T) {
	runSample(t, `500000
5
196303 288435 1
181743 468082 1
454772 467304 2
14914 392969 1
253044 366728 1
`, []int{407867, 213660, 226193, 59364, 59364})
}
