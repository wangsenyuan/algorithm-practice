package p3568

import "testing"

func runSample(t *testing.T, classroom []string, energy int, expect int) {
	res := minMoves(classroom, energy)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	classroom := []string{"S.", "XL"}
	energy := 2
	expect := 2
	runSample(t, classroom, energy, expect)
}

func TestSample2(t *testing.T) {
	classroom := []string{"LS", "RL"}
	energy := 4
	expect := 3
	runSample(t, classroom, energy, expect)
}

func TestSample3(t *testing.T) {
	classroom := []string{"L.S", "RXL"}
	energy := 3
	expect := -1
	runSample(t, classroom, energy, expect)
}
