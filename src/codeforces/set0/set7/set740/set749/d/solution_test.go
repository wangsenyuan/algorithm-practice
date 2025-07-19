package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, ans := range res {
		expect := readNNums(reader, 2)
		if !reflect.DeepEqual(ans, expect) {
			t.Fatalf("Sample expect %v, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
1 10
2 100
3 1000
1 10000
2 100000
3 1000000
3
1 3
2 2 3
2 1 2
2 100000
1 10
3 1000`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 10
2 100
1 1000
2
2 1 2
2 2 3
0 0
1 10`)
}
