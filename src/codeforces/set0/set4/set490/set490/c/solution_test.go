package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	x, a, b, ok, s1, s2 := drive(bufio.NewReader(strings.NewReader(s)))

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}

	if s1+s2 != x {
		t.Fatalf("Sample result %s + %s, not %s", s1, s2, x)
	}
	var num1 int
	for i := 0; i < len(s1); i++ {
		num1 = add(mul(num1, 10, a), int(s1[i]-'0'), a)
	}
	var num2 int
	for i := 0; i < len(s2); i++ {
		num2 = add(mul(num2, 10, b), int(s2[i]-'0'), b)
	}

	if num1 != 0 || num2 != 0 {
		t.Fatalf("Sample result %s => %d, %s => %d", s1, num1, s2, num2)
	}
}

func TestSample1(t *testing.T) {
	s := `116401024
97 1024
`
	expect := true

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `284254589153928171911281811000
1009 1000
`
	expect := true

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `120
12 1
`
	expect := false

	runSample(t, s, expect)
}
