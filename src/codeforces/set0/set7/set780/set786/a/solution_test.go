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
	expect := []string{
		readString(reader),
		readString(reader),
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
2 3 2
3 1 2 3
Lose Win Win Loop
Loop Win Win Win`)
}

func TestSample2(t *testing.T) {
	runSample(t, `8
4 6 2 3 4
2 3 6
Win Win Win Win Win Win Win
Lose Win Lose Lose Win Lose Lose`)
}
