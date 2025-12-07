package p3771

import "testing"

func runSample(t *testing.T, hp int, damage []int, requirement []int, expect int64) {
	res := totalScore(hp, damage, requirement)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	hp := 11
	damage := []int{3, 6, 7}
	requirement := []int{4, 2, 5}
	expect := int64(3)
	runSample(t, hp, damage, requirement, expect)
}

func TestSample2(t *testing.T) {
	hp := 2
	damage := []int{10000, 1}
	requirement := []int{1, 1}
	expect := int64(1)
	runSample(t, hp, damage, requirement, expect)
}
