package p3579

import "testing"

func runSample(t *testing.T, w1 string, w2 string, expect int) {
	res := minOperations(w1, w2)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcdf", "dacbe", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "abceded", "baecfef", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "abc", "bca", 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "abcdef", "fedabc", 2)
}
