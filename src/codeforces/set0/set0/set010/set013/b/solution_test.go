package main

import "testing"

func runSample(t *testing.T, seg1, seg2, seg3 Segment, expect bool) {
	res := solve(seg1, seg2, seg3)

	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// Example 1: Should return YES
	seg1 := Segment{Point{4, 4}, Point{6, 0}}
	seg2 := Segment{Point{4, 1}, Point{5, 2}}
	seg3 := Segment{Point{4, 0}, Point{4, 4}}
	runSample(t, seg1, seg2, seg3, true)
}

func TestSample2(t *testing.T) {
	// Example 2: Should return NO
	seg1 := Segment{Point{0, 0}, Point{0, 6}}
	seg2 := Segment{Point{0, 6}, Point{2, -4}}
	seg3 := Segment{Point{1, 1}, Point{0, 1}}
	runSample(t, seg1, seg2, seg3, false)
}

func TestSample3(t *testing.T) {
	// Example 3: Should return YES
	seg1 := Segment{Point{0, 0}, Point{0, 5}}
	seg2 := Segment{Point{0, 5}, Point{2, -1}}
	seg3 := Segment{Point{1, 2}, Point{0, 1}}
	runSample(t, seg1, seg2, seg3, true)
}
