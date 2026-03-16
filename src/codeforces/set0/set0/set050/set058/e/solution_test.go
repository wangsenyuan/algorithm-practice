package main

import (
	"testing"
)

func runSample(t *testing.T, expr string, expect string) {
	res := solve(expr)
	if len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	a1, b1, c1 := parseExpr(expr)
	a2, b2, c2 := parseExpr(res)

	contains := func(s1 string, s2 string) bool {
		for i, j := 0, 0; i < len(s2); i++ {
			for j < len(s1) && s1[j] != s2[i] {
				j++
			}
			if j == len(s1) {
				return false
			}
			j++
		}
		return true
	}

	if !contains(a2, a1) || !contains(b2, b1) || !contains(c2, c1) || len(c2) < max(len(a2), len(b2)) {
		t.Fatalf("Sample result %s, not valid", res)
	}

	// a2 + b2 == c2
	var carry int
	for i := 0; i < len(c2); i++ {
		var sum int
		if i < len(a2) {
			sum += int(a2[i] - '0')
		}
		if i < len(b2) {
			sum += int(b2[i] - '0')
		}
		sum += carry
		carry = sum / 10
		x := sum % 10
		if x != int(c2[i]-'0') {
			t.Fatalf("Sample result %s, not valid", res)
		}
	}
	if carry != 0 {
		t.Fatalf("Sample result %s, not valid", res)
	}
}

func TestSample1(t *testing.T) {
	expr := "2+4=5"
	expect := "21+4=25"
	runSample(t, expr, expect)
}

func TestSample2(t *testing.T) {
	expr := "1+1=3"
	expect := "1+31=32"
	runSample(t, expr, expect)
}

func TestSample3(t *testing.T) {
	expr := "1+1=2"
	expect := "1+1=2"
	runSample(t, expr, expect)
}

func TestSample4(t *testing.T) {
	expr := "123122+765654=975632"
	expect := "123192200+76565432=199757632"
	runSample(t, expr, expect)
}

func TestSample5(t *testing.T) {
	expr := "323+123=287"
	expect := "323+12553=12876"
	runSample(t, expr, expect)
}

func TestSample6(t *testing.T) {
	expr := "100+10=96454"
	expect := "86100+10354=96454"
	runSample(t, expr, expect)
}
