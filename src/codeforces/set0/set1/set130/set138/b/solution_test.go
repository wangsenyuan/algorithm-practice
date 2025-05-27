package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	num, s1, s2 := process(reader)
	x1 := readString(reader)
	x2 := readString(reader)

	count := func(s string) []int {
		res := make([]int, 10)
		for _, x := range []byte(s) {
			res[int(x-'0')]++
		}
		return res
	}
	cnt := count(num)

	checkCount := func(s string) {
		tmp := count(s)
		if !reflect.DeepEqual(tmp, cnt) {
			t.Fatalf("string %s is not a permuation of %s", s, num)
		}
	}
	n := len(num)
	check := func(a string, b string) int {
		checkCount(a)
		checkCount(b)
		var res int
		var carry int
		for i := n - 1; i >= 0; i-- {
			x := int(a[i] - '0')
			y := int(b[i] - '0')
			if (x+y+carry)%10 == 0 {
				res++
				carry = (x + y + carry) / 10
			} else {
				break
			}
		}
		return res
	}

	a := check(x1, x2)
	b := check(s1, s2)
	if a != b {
		t.Fatalf("Sample result %s %s, not correct", s1, s2)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `198
981
819
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `500
500
500
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `1099
9901
1099
`)
}


func TestSample4(t *testing.T) {
	runSample(t, `1061
6110
6110
`)
}
