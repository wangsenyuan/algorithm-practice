package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	expect_cnt := readNum(reader)
	expect_res := readNNums(reader, len(res))

	if cnt != expect_cnt || !reflect.DeepEqual(res, expect_res) {
		t.Fatalf("Sample expect %d, %v, but got %d %v", expect_cnt, expect_res, cnt, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 1
5 2
10 1
5
0 1
1 3
3 0
4 0
4 0
`
	expect := `2
3 3 -1 
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
3 2
7 1
11 2
4
2 1
6 0
6 4
11 2
`
	expect := `3
1 2 4 
`
	runSample(t, s, expect)
}
