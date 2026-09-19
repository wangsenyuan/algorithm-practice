package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
Troll likes Dracul
Dracul likes Anka
Snowy likes Hexadecimal
210 200 180
`
	expect := []int{30, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
Anka likes Chapay
Chapay likes Anka
10000 50 50
`
	expect := []int{1950, 2}
	runSample(t, s, expect)
}
