package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
AbRb r Zz
4
xR abRb
aA xr
zz Z
xr y`
	expect := []int{2, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
RuruRu fedya
1
ruruRU fedor
`
	expect := []int{1, 10}
	runSample(t, s, expect)
}
