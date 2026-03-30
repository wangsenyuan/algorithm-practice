package p3883

import "testing"

func runSample(t *testing.T, digitSum []int, expect int) {
	res := countArrays(digitSum)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	digitSum := []int{25, 1}
	expect := 6
	runSample(t, digitSum, expect)
}

func TestSample2(t *testing.T) {
	digitSum := []int{1}
	expect := 4
	runSample(t, digitSum, expect)
}

func TestSample3(t *testing.T) {
	digitSum := []int{2, 49, 23}
	expect := 0
	runSample(t, digitSum, expect)
}
