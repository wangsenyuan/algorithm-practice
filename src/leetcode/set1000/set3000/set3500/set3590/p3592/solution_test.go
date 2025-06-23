package p3592

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, numWays []int, expect []int) {
	res := findCoins(numWays)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	numWays := []int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}
	expect := []int{2, 4, 6}
	runSample(t, numWays, expect)
}

func TestSample2(t *testing.T) {
	numWays := []int{1, 2, 2, 3, 4}
	expect := []int{1, 2, 5}
	runSample(t, numWays, expect)
}

func TestSample3(t *testing.T) {
	numWays := []int{1, 2, 3, 4, 15}
	runSample(t, numWays, nil)
}
