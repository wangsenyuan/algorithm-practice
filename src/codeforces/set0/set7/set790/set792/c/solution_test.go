package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, num := process(reader)

	expect := readString(reader)

	if res == expect {
		return
	}
	if expect == "-1" || res == "-1" || len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	if len(res) == 1 {
		x := int(res[0] - '0')
		if x%3 != 0 {
			t.Fatalf("Sample expect %s, but got %s", expect, res)
		}
		return
	}
	if res[0] == '0' {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	var sum int
	for i, j := 0, 0; i < len(res); i++ {
		for j < len(num) && num[j] != res[i] {
			j++
		}
		if j == len(num) {
			t.Fatalf("Sample expect %s, but got %s", expect, res)
		}
		sum += int(res[i] - '0')
		j++
	}
	if sum%3 != 0 {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1033
33`)
}

func TestSample2(t *testing.T) {
	runSample(t, `10
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `11
-1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `20000111
200001`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5388306043547446322173224045662327678394712363272776811399689704247387317165308057863239568137902157
538830603547446322173224045662327678394712363272776811399689704247387317165308057863239568137902157`)
}

func TestSample6(t *testing.T) {
	runSample(t, `4902501252475186372406731932548506197390793597574544727433297197476846519276598727359617092494798814
490501252475186372406731932548506197390793597574544727433297197476846519276598727359617092494798814`)
}
