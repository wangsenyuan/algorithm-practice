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
	m := readNum(reader)
	if len(res) != m {
		t.Fatalf("Sample expect %d instructions, but got %v", m, res)
	}
	if m == 0 {
		return
	}
	ans := readNNums(reader, m)
	if !reflect.DeepEqual(res, ans) {
		t.Fatalf("Sample expect %v, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 4
3 3
3 1
4 1
9 2
2
2 3 
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 1
1 1
0
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `20 10
13 3
4 3
13 5
1 2
1 5
5 2
11 2
5 2
4 4
12 4
6
1 2 4 6 8 10
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `20 10
7 2
12 2
14 1
4 3
17 4
9 3
15 4
12 2
18 1
12 1
3
8 9 10 
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `20 10
7 2
8 1
1 3
13 2
13 1
13 2
2 3
3 2
15 3
3 1
5
2 5 6 8 10 
`
	runSample(t, s)
}
