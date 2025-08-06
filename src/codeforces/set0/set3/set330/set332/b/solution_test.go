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
	ans := process(reader)
	expect := make([]int, 2)
	fmt.Fscan(reader, &expect[0], &expect[1])
	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 2
3 6 1 1 6
1 4`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6 2
1 1 1 1 1 1
1 3`)
}
