package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	expect := make([]int, len(res))
	for i := range expect {
		fmt.Fscan(reader, &expect[i])
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 4
4 3 1 6
3 2 5 4`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
3 2 0 2 7
2 1 4 1 6`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
2 3 1
1 2 3`)
}
