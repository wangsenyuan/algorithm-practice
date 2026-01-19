package p3816

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := lexSmallestAfterDeletion(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aaccb", "aacb")
}

func TestSample2(t *testing.T) {
	runSample(t, "z", "z")
}

func TestSample3(t *testing.T) {
	runSample(t, "bbcac", "bac")
}


func TestSample4(t *testing.T) {
	runSample(t, "aaccbdfefeadfafefafdfe", "aacbaaadfe")
}