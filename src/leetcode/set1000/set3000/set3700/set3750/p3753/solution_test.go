package p3753

import "testing"

func runSample(t *testing.T, num1 int, num2 int, expect int) {
	res := totalWaviness(int64(num1), int64(num2))
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num1, num2 := 120, 130
	expect := 3
	runSample(t, num1, num2, expect)
}

func TestSample2(t *testing.T) {
	num1, num2 := 198, 202
	expect := 3
	runSample(t, num1, num2, expect)
}

func TestSample3(t *testing.T) {
	num1, num2 := 4848, 4848
	expect := 2
	runSample(t, num1, num2, expect)
}