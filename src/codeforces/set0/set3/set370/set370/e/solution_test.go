package main

import "testing"

func validate(t *testing.T, input []int, best int, res []int) {
	t.Helper()
	if best < 0 {
		if res != nil {
			t.Fatalf("expected nil result for impossible case")
		}
		return
	}
	if len(res) != len(input) {
		t.Fatalf("length mismatch: expect %d, got %d", len(input), len(res))
	}
	for i, x := range input {
		if x != 0 && res[i] != x {
			t.Fatalf("position %d: expect fixed value %d, got %d", i, x, res[i])
		}
	}
	cur := 0
	books := 0
	for cur < len(res) {
		books++
		if res[cur] != books {
			t.Fatalf("book numbering broken at day %d: expect %d, got %d", cur, books, res[cur])
		}
		j := cur
		for j < len(res) && res[j] == books {
			j++
		}
		if j-cur < 2 || j-cur > 5 {
			t.Fatalf("book %d has invalid length %d", books, j-cur)
		}
		cur = j
	}
	if books != best {
		t.Fatalf("expect %d books, got %d", best, books)
	}
}

func runSample(t *testing.T, input []int, expect int) {
	t.Helper()
	best, res := solve(input)
	if best != expect {
		t.Fatalf("expect %d books, got %d", expect, best)
	}
	validate(t, input, best, res)
}

func TestSample1(t *testing.T) {
	runSample(t, []int{0, 1, 0, 0, 0, 3, 0}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{0, 0, 0, 0, 0, 0, 0, 0}, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 0, 1, 0}, 1)
}

func TestSample4(t *testing.T) {
	best, res := solve([]int{0, 0, 0, 3})
	if best != -1 || res != nil {
		t.Fatalf("expect impossible, got %d %v", best, res)
	}
}

func TestImpossibleDecreasing(t *testing.T) {
	best, res := solve([]int{1, 0, 2, 0, 1, 0})
	if best != -1 || res != nil {
		t.Fatalf("expect impossible, got %d %v", best, res)
	}
}

func TestObservedBookNeedsLongerSegment(t *testing.T) {
	runSample(t, []int{0, 0, 2, 0, 2, 0, 3, 0, 0}, 4)
}

func TestImpossibleOneDaySuffix(t *testing.T) {
	best, res := solve([]int{1, 1, 1, 1, 1, 0})
	if best != -1 || res != nil {
		t.Fatalf("expect impossible, got %d %v", best, res)
	}
}

func TestObservedBookIndexTooLarge(t *testing.T) {
	best, res := solve([]int{1, 3})
	if best != -1 || res != nil {
		t.Fatalf("expect impossible, got %d %v", best, res)
	}
}

func TestBlockCanReachLengthTwoFromLeft(t *testing.T) {
	runSample(t, []int{0, 0, 0, 2, 0, 3}, 3)
}
