package p3767

import "testing"

func runSample(t *testing.T, technique1 []int, technique2 []int, k int, expect int64) {
	res := maxPoints(technique1, technique2, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	technique1 := []int{5, 2, 10}
	technique2 := []int{10, 3, 8}
	k := 2
	expect := int64(22)
	runSample(t, technique1, technique2, k, expect)
}

func TestSample2(t *testing.T) {
	technique1 := []int{10, 20, 30}
	technique2 := []int{5, 15, 25}
	k := 2
	expect := int64(60)
	runSample(t, technique1, technique2, k, expect)
}

func TestSample3(t *testing.T) {
	technique1 := []int{1, 2, 3}
	technique2 := []int{4, 5, 6}
	k := 0
	expect := int64(15)
	runSample(t, technique1, technique2, k, expect)
}
