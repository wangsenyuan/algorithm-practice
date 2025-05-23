package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	expect := readNNums(reader, 2)
	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 2 3
3
1 2
2 3
3 2
3 1
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
2 8 0 6 0
6
1 4
1 3
2 4
3 4
4 5
5 1
8 2
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
1 3 2 2 1 3 1 4 10 10
12
1 2
2 3
3 1
3 4
4 5
5 6
5 7
6 4
7 3
8 9
9 10
10 9
15 6
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `2
7 91
2
1 2
2 1
7 1
	`)
}
