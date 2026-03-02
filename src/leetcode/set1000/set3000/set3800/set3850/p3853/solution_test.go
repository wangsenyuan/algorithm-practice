package p3853

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := mergeCharacters(s, k)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "yybyzybz", 2, "ybzybz")
}
