package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	tmp := output(ans)

	if tmp != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, tmp)
	}
}

func TestSample1(t *testing.T) {
	s := `2
5
-> 4 3
<- 15
<- 4
-> 3 1
-> 1 3`
	expect := `7
2 3
1 2
9
13
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	expect := `1
4
3
2
1 1
2 2
2 1
1 2
`
	s := `1
8
-> 1 1
-> 1 2
-> 2 1
-> 2 2
<- 1
<- 2
<- 3
<- 4
`
	runSample(t, s, expect)
}
