package main

import (
	"slices"
	"strconv"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if len(res) != expect {
		t.Fatalf("expect %d, but got %d", expect, len(res))
	}

	if res[0][0] != '+' {
		t.Fatalf("Sample result %v, not correct", res)
	}

	m, _ := strconv.Atoi(res[0][3:])
	a := make([]int, m+1)
	a[m] = 1
	for i := 1; i < len(res); i++ {
		j, _ := strconv.Atoi(res[i][3:])
		if res[i][0] == '+' {
			c := 1
			for j <= m && c > 0 {
				v := a[j] + c
				a[j] = v % 2
				c = v / 2
				j++
			}
			if c == 1 {
				t.Fatalf("Sample result %v, not correct", res)
			}
		} else {
			k := j
			for k <= m && a[k] == 0 {
				a[k] = 1
				k++
			}
			if k > m {
				t.Fatalf("Sample result %v, not correct", res)
			}
			// 要进行减操作, a[j]必须是1，而且是要通过借位产生
			a[k] = 0
		}
	}

	buf := []byte(s)
	slices.Reverse(buf)

	for i := 0; i < len(buf) && i < len(a); i++ {
		x := int(buf[i] - '0')
		if x != a[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	if len(buf) < len(a) && a[m] == 1 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "1111"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1010011"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "10110111"
	expect := 4
	runSample(t, s, expect)
}
