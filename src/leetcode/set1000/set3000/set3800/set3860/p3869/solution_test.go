package p3869

import "testing"

func runSample(t *testing.T, l int64, r int64, expect int64) {
	res := countFancy(l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 10, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 12340, 12341, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 123456788, 123456788, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 123456788, 112233581)
}

func TestSample5(t *testing.T) {
	runSample(t, 144728015978, 146069087503, 1219155934)
}

func TestSample6(t *testing.T) {
	runSample(t, 197957117508, 319024461148, 110061221477)
}
