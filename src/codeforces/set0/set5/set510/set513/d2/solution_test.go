package main

import (
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
1 2 LEFT
1 3 RIGHT
`, []int{2, 1, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 2 RIGHT
1 3 LEFT
`, nil)
}

func TestContradictoryConstraints(t *testing.T) {
	runSample(t, `2 2
1 2 LEFT
1 2 RIGHT
`, nil)
}

func generateInorders(root int, size int) [][]int {
	if size == 0 {
		return [][]int{{}}
	}
	var res [][]int
	for leftSize := 0; leftSize < size; leftSize++ {
		rightSize := size - leftSize - 1
		left := generateInorders(root+1, leftSize)
		right := generateInorders(root+leftSize+1, rightSize)
		for _, a := range left {
			for _, b := range right {
				cur := make([]int, 0, size)
				cur = append(cur, a...)
				cur = append(cur, root)
				cur = append(cur, b...)
				res = append(res, cur)
			}
		}
	}
	return res
}

func satisfiesConstraints(inorder []int, constraints []constraint) bool {
	n := len(inorder)
	leftEnd := make([]int, n+1)
	subtreeEnd := make([]int, n+1)

	var build func(int, []int) bool
	build = func(root int, order []int) bool {
		if len(order) == 0 {
			return true
		}
		pos := -1
		for i, label := range order {
			if label == root {
				pos = i
				break
			}
		}
		if pos < 0 {
			return false
		}

		leftEnd[root] = root + pos
		subtreeEnd[root] = root + len(order) - 1
		return build(root+1, order[:pos]) &&
			build(root+pos+1, order[pos+1:])
	}

	if !build(1, inorder) {
		return false
	}
	for _, cur := range constraints {
		if cur.direction == "LEFT" {
			if !(cur.a < cur.b && cur.b <= leftEnd[cur.a]) {
				return false
			}
		} else if !(leftEnd[cur.a] < cur.b && cur.b <= subtreeEnd[cur.a]) {
			return false
		}
	}
	return true
}

func TestSolveAgainstExhaustiveTrees(t *testing.T) {
	rng := rand.New(rand.NewSource(5132))
	for n := 1; n <= 7; n++ {
		trees := generateInorders(1, n)
		for test := 0; test < 1000; test++ {
			c := rng.Intn(11)
			constraints := make([]constraint, c)
			for i := range constraints {
				constraints[i].a = rng.Intn(n) + 1
				constraints[i].b = rng.Intn(n) + 1
				if rng.Intn(2) == 0 {
					constraints[i].direction = "LEFT"
				} else {
					constraints[i].direction = "RIGHT"
				}
			}

			possible := false
			for _, inorder := range trees {
				if satisfiesConstraints(inorder, constraints) {
					possible = true
					break
				}
			}

			result := solve(n, constraints)
			if possible != (result != nil) {
				t.Fatalf("n = %d, constraints = %v, possible = %t, result = %v",
					n, constraints, possible, result)
			}
			if result != nil && !satisfiesConstraints(result, constraints) {
				t.Fatalf("n = %d, constraints = %v, invalid result = %v",
					n, constraints, result)
			}
		}
	}
}

func TestMaximumN(t *testing.T) {
	const n = 1_000_000
	result := solve(n, []constraint{{a: 1, b: n, direction: "RIGHT"}})
	if len(result) != n {
		t.Fatalf("expect %d nodes, got %d", n, len(result))
	}
	for i, label := range result {
		if label != i+1 {
			t.Fatalf("position %d: expect %d, got %d", i, i+1, label)
		}
	}
}
