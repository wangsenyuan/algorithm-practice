package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	var got string
	if len(res) == 0 {
		got = "No solution"
	} else {
		tmp := fmt.Sprintf("%v", res)
		got = tmp[1 : len(tmp)-1]
	}
	if got != expect {
		t.Errorf("Sample expect %s, but got %s", expect, got)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 3 1`
	expect := "1 3 2 1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 3 2 1`
	expect := "No solution"
	runSample(t, s, expect)
}
