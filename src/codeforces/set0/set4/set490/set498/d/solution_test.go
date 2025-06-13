package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	m := readNum(reader)
	if len(ans) != m {
		t.Fatalf("Samplee expect %d, but got %v", m, ans)
	}
	for i := range m {
		x := readNum(reader)
		if x != ans[i] {
			t.Fatalf("Sample expect %d, but got %d", x, ans[i])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `10
2 5 3 2 3 5 3 4 2 4
10
C 10 6
A 2 6
A 1 3
C 3 4
A 3 11
A 4 9
A 5 6
C 7 3
A 8 10
A 2 5
7
5
3
14
6
2
4
4
`
	runSample(t, s)
}


func TestSample2(t *testing.T) {
	s := `10
2 5 4 2 3 5 3 4 2 6
1
A 3 11
1
14
`
	runSample(t, s)
}