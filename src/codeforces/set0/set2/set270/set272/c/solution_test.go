package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	for _, x := range ans {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Errorf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2 3 6 6
4
1 1
3 1
1 1
4 3
1
3
4
6`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 3
2
1 1
3 1
1
3`)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
1
5
1 2
1 10
1 10
1 10
1 10
1
3
13
23
33`)
}
