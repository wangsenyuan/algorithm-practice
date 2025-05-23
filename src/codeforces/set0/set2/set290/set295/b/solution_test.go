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
	expect := readNNums(reader, len(ans))
	if !reflect.DeepEqual(ans, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
0
1
0
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
0 5
4 0
1 2
9 0
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
0 3 1 1
6 0 400 1
2 4 0 1
1 1 1 0
4 1 2 3
17 23 404 0 
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `5
0 27799 15529 16434 44291
47134 0 90227 26873 52252
41605 21269 0 9135 55784
70744 17563 79061 0 73981
70529 35681 91073 52031 0
5 2 3 1 4
896203 429762 232508 87178 0 
	`)
}
