package p3746

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := minLengthAfterRemovals(s)
	if res != expect {
		t.Errorf("Sample failed, expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `aabbab`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `aaaa`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `aaabb`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `babba`
	expect := 1
	runSample(t, s, expect)
}
