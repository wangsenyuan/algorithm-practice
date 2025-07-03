package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ok, s1, s2, nums := process(reader)
	expect := readString(reader)
	if expect == "Happy new year!" {
		if ok {
			t.Fatalf("Sample expect %s, but got %t", expect, ok)
		}
		return
	}

	if !ok {
		t.Fatalf("Sample expect %s, but got %t", expect, ok)
	}

	if len(s1) != nums[2] || len(s2) != nums[3] {
		t.Fatalf("Sample result %s %s, not correct", s1, s2)
	}

	type data struct {
		cnt   int
		first byte
		last  byte
	}

	parse := func(x string) data {
		var cnt int
		for i := 0; i+1 < len(x); i++ {
			if x[i] == 'A' && x[i+1] == 'C' {
				cnt++
			}
		}
		return data{cnt, x[0], x[len(x)-1]}
	}

	d1 := parse(s1)
	d2 := parse(s2)

	for i := 3; i <= nums[0]; i++ {
		cnt := d1.cnt + d2.cnt
		if d1.last == 'A' && d2.first == 'C' {
			cnt++
		}
		d := data{cnt, d1.first, d2.last}
		d1, d2 = d2, d
	}
	if d2.cnt != nums[1] {
		t.Fatalf("Sample result %s %s, not correct", s1, s2)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2 2 2
AC
AC
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3 2 2
Happy new year!
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3 2 1
Happy new year!
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 2 2 1
Happy new year!
`)
}
