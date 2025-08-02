package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, v := range res {
		var u int
		fmt.Fscan(reader, &u)
		if u != v {
			t.Fatalf("Sample expect %d, but got %d", v, u)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
0 1 2 3 4
4 3
2 0
3 7
5 7
5 8
4
2
0
4
0	`)
}


func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 1 1
3 1
2 0
4
2`)
}
