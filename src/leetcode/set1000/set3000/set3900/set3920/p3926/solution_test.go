package p3926

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, chunks []string, queries []string, expect []int) {
	res := countWordOccurrences(chunks, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	chunks := []string{"hello wor", "ld hello"}
	queries := []string{"hello", "world", "wor"}
	expect := []int{2, 1, 0}
	runSample(t, chunks, queries, expect)
}

func TestSample2(t *testing.T) {
	chunks := []string{"a--b a-", "-c"}
	queries := []string{"a", "b", "c"}
	expect := []int{2, 1, 1}
	runSample(t, chunks, queries, expect)
}

func TestSample3(t *testing.T) {
	chunks := []string{"hello"}
	queries := []string{"hello", "ell"}
	expect := []int{1, 0}
	runSample(t, chunks, queries, expect)
}
