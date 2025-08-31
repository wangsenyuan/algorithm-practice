package p3669

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := minDifference(n, k)
	sort.Ints(expect)
	sort.Ints(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 180
	k := 2
	expect := []int{12, 15}
	runSample(t, n, k, expect)
}
