package p3906

import "testing"

func runSample(t *testing.T, l int64, r int64, directions string, expect int64) {
	res := countGoodIntegersOnPath(l, r, directions)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 10, "DDDRRR", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 123456789, 123456790, "DDRRDR", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1288561398769758, 1288561398769758, "RRRDDD", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 100, 1288561398769758, "RRRDDD", 7080243999900)
}
