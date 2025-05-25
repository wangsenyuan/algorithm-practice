package p3563

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := lexicographicallySmallestString(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abc", "a")
}
func TestSample2(t *testing.T) {
	runSample(t, "bcda", "")
}

func TestSample3(t *testing.T) {
	runSample(t, "zdce", "zdce")
}
