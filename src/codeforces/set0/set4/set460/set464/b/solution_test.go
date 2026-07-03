package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res.ok != expect {
		t.Fatalf("Sample ok expect %t, but got %t", expect, res.ok)
	}
	if !expect {
		return
	}

	if !check(res.pts) {
		t.Fatalf("Sample pts expect to be a cube, but got %v", res.pts)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0 0 0
0 0 1
0 0 1
0 0 1
0 1 1
0 1 1
0 1 1
1 1 1
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `0 0 0
0 0 0
0 0 0
0 0 0
1 1 1
1 1 1
1 1 1
1 1 1
`, false)
}

func TestCheckCubeAcceptsAxisAlignedCube(t *testing.T) {
	pts := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}

	if !check(pts) {
		t.Fatalf("check expect true for axis-aligned cube, but got false")
	}
}

func TestCheckCubeAcceptsRotatedCube(t *testing.T) {
	u := []int{1, 2, 2}
	v := []int{2, 1, -2}
	w := []int{-2, 2, -1}
	pts := make([][]int, 0, 8)
	for mask := range 8 {
		p := []int{0, 0, 0}
		for j, vec := range [][]int{u, v, w} {
			if mask>>j&1 == 1 {
				for k := range 3 {
					p[k] += vec[k]
				}
			}
		}
		pts = append(pts, p)
	}

	if !check(pts) {
		t.Fatalf("check expect true for rotated cube, but got false for %v", pts)
	}
}

func TestCheckCubeRejectsNonCubeBox(t *testing.T) {
	pts := [][]int{
		{0, 0, 0},
		{2, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{2, 1, 0},
		{2, 0, 1},
		{0, 1, 1},
		{2, 1, 1},
	}

	if check(pts) {
		t.Fatalf("check expect false for rectangular box with unequal side lengths")
	}
}

func TestCheckCubeRejectsDuplicatePoints(t *testing.T) {
	pts := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 1},
	}

	if check(pts) {
		t.Fatalf("check expect false when points contain duplicates")
	}
}

func TestNextPermutationAdvancesLexicographically(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := [][]int{
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, cur := range expect {
		got := nextPermutation(nums)
		if !reflect.DeepEqual(got, cur) {
			t.Fatalf("nextPermutation expect %v, but got %v", cur, got)
		}
		if !reflect.DeepEqual(nums, cur) {
			t.Fatalf("nextPermutation should mutate nums to %v, but got %v", cur, nums)
		}
	}
}

func TestNextPermutationSkipsDuplicatePermutations(t *testing.T) {
	nums := []int{0, 0, 1}
	expect := [][]int{
		{0, 1, 0},
		{1, 0, 0},
	}

	for _, cur := range expect {
		got := nextPermutation(nums)
		if !reflect.DeepEqual(got, cur) {
			t.Fatalf("nextPermutation expect %v, but got %v", cur, got)
		}
	}

	if got := nextPermutation(nums); got != nil {
		t.Fatalf("nextPermutation expect nil after last unique permutation, but got %v", got)
	}
}

func TestNextPermutationReturnsNilForDescendingOrder(t *testing.T) {
	nums := []int{3, 2, 1}
	if got := nextPermutation(nums); got != nil {
		t.Fatalf("nextPermutation expect nil for final permutation, but got %v", got)
	}
	if expect := []int{3, 2, 1}; !reflect.DeepEqual(nums, expect) {
		t.Fatalf("nextPermutation should keep final permutation as %v, but got %v", expect, nums)
	}
}
