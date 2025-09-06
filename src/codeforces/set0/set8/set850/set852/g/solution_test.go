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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1
abc
aec
ac
a?c
`
	expect := []int{3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `22 2
aaaab
aaabb
aabab
aabbb
abaab
ababb
abbab
abbbb
aaab
aabb
abab
abbb
aab
abb
ab
cccd
ccdd
cdcd
cddd
ccd
cdd
cd
a???b
c??d
`
	expect := []int{15, 7}
	runSample(t, s, expect)
}
