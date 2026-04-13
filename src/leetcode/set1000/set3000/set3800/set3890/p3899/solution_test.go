package p3899

import (
	"testing"
)

func runSample(t *testing.T, sides []int, expect bool) {
	res := internalAngles(sides)

	if len(res) == 3 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	sides := []int{3, 4, 5}
	runSample(t, sides, true)
}
