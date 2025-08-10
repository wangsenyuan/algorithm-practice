package p3646

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := specialPalindrome(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 22)
}

func TestSample2(t *testing.T) {
	runSample(t, 33, 212)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 212)
}

func TestSample4(t *testing.T) {
	runSample(t, 100401023, 234434432)
}

func TestSample5(t *testing.T) {
	runSample(t, 210, 212)
}

func TestSample6(t *testing.T) {
	runSample(t, 216, 333)
}

func TestSample7(t *testing.T) {
	runSample(t, 916681587, 999999999)
}
