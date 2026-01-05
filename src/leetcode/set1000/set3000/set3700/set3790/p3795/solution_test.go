package p3795

import "testing"

func runSample(t *testing.T, grid []string, d int, expect int) {
	res := numberOfRoutes(grid, d)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"..",
		"#.",
	}
	d := 1
	expect := 2
	runSample(t, grid, d, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"..",
		"#.",
	}
	d := 2
	expect := 4
	runSample(t, grid, d, expect)
}


func TestSample3(t *testing.T) {
	grid := []string{
		"..",
	}
	d := 1
	expect := 4
	runSample(t, grid, d, expect)
}