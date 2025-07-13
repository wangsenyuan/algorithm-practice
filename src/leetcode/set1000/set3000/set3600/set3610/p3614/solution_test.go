package p3614

import "testing"

func runSample(t *testing.T, s string, K int64, expect byte) {
	res := processStr(s, K)
	if res != expect {
		t.Errorf("Sample expect %c, but got %c", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "a#b%*", 1, 'a')
}

func TestSample2(t *testing.T) {
	runSample(t, "cd%#*#", 3, 'd')
}

func TestSample3(t *testing.T) {
	runSample(t, "z*#", 0, '.')
}
