package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	for _, y := range ans {
		x := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %v", s, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6
A
B
BA
BB
AA
AB
0
1
2
5
10
25
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10
A
B
AABA
AABB
AB
AA
AAA
BB
AAB
BA
0
1
2
4
8
20
41
82
170
425
`
	runSample(t, s)
}
