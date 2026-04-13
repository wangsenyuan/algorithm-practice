package p3901

import "testing"

func runSample(t *testing.T, nums []int, p int, queries [][]int, expect int) {
	res := countGoodSubseq(nums, p, queries)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 8, 12, 16}
	p := 2
	queries := [][]int{{0, 3}, {2, 6}}
	expect := 1
	runSample(t, nums, p, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 5, 7, 8}
	p := 3
	queries := [][]int{{0, 6}, {1, 9}, {2, 3}}
	expect := 2
	runSample(t, nums, p, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{14, 4, 13, 21}
	p := 3
	queries := [][]int{{0, 8}, {0, 12}}
	expect := 1
	runSample(t, nums, p, queries, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{21, 64, 95}
	p := 1
	queries := [][]int{
		{0, 56}, {2, 96}, {2, 49}, {1, 18}, {0, 36}, {0, 82}, {2, 56}, {1, 79}, {1, 24},
		{2, 95}, {2, 27}, {2, 12}, {0, 79}, {1, 54}, {1, 41}, {1, 66}, {2, 36}, {2, 82},
		{0, 100}, {0, 95}, {1, 1}, {2, 24}, {1, 93}, {0, 32}, {0, 64}, {2, 47}, {2, 36},
		{0, 11}, {0, 66}, {2, 13}, {2, 38}, {1, 57}, {1, 65}, {0, 15}, {0, 21}, {1, 64},
		{2, 95},
	}
	expect := 30
	runSample(t, nums, p, queries, expect)
}
