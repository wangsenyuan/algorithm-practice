package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
+*
*U
1 1
2 2`
	expect := -1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
<><
><>
1 1
2 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 10
>+^+U-DU>*
ULLL*UL+>+
U<>>L^D>>v
|*L+^^R^R^
LD+|L*<D*D
+>U^|UL+-R
D>vvR+R|^D
*+v^><^vLL
LU^|^U->D|
*D>-|>+^L>
3 8
4 5`
	expect := 14
	runSample(t, s, expect)
}
