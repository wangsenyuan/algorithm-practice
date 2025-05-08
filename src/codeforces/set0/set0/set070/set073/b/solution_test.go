package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
teama 10
teamb 20
teamc 40
2
10 20
teama
`
	expect := []int{2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
teama 10
teamb 10
2
10 10
teamb
`
	expect := []int{2, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20
g 1000000
ibb 1000000
idj 1000000
kccg 1000000
bbe 1000000
hjf 1000000
a 1000000
f 1000000
ijj 1000000
akgf 1000000
kdkhj 1000000
e 1000000
h 1000000
hb 1000000
faie 1000000
j 1000000
i 1000000
hgg 1000000
fi 1000000
icf 1000000
12
1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000 1000000
a
`
	expect := []int{1, 13}
	runSample(t, s, expect)
}
