package p3765

import "testing"

func runSample(t *testing.T, num int, expect bool) {
	res := completePrime(num)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 23
	expect := true
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := 39
	expect := false
	runSample(t, num, expect)
}
