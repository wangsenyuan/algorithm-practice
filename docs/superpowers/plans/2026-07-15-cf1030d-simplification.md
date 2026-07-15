# Codeforces 1030D Simplification Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the accepted divisor-enumeration construction with the standard two-GCD construction, strengthen its tests, and document the final algorithm.

**Architecture:** Keep the existing axis-aligned triangle and package interfaces. Reduce `k` against `n` and `m`; the unreduced factor must be at most two, and a remaining factor of one is handled by doubling one reduced side that still fits.

**Tech Stack:** Go standard library, package-local `go test`, Markdown.

---

### Task 1: Add exhaustive characterization coverage

**Files:**
- Modify: `src/codeforces/set1/set10/set103/set1030/d/solution_test.go`

- [ ] **Step 1: Remove the duplicate sample**

Delete `TestSample3`, which repeats input `4 4 7` and expected result `false`.

- [ ] **Step 2: Add an exhaustive small-input test**

Add this test after the remaining samples:

```go
func TestSmall(t *testing.T) {
	for n := 1; n <= 8; n++ {
		for m := 1; m <= 8; m++ {
			for k := 2; k <= 2*n*m+2; k++ {
				res := solve(n, m, k)
				expectOK := (2*n*m)%k == 0
				if res.ok != expectOK {
					t.Fatalf("n=%d m=%d k=%d: expect ok=%v, got %v", n, m, k, expectOK, res)
				}
				if !res.ok {
					continue
				}
				if len(res.pts) != 3 {
					t.Fatalf("n=%d m=%d k=%d: expect 3 points, got %v", n, m, k, res.pts)
				}
				for _, p := range res.pts {
					if p[0] < 0 || p[0] > n || p[1] < 0 || p[1] > m {
						t.Fatalf("n=%d m=%d k=%d: point %v is out of bounds", n, m, k, p)
					}
				}
				if area2(res.pts) != 2*n*m/k {
					t.Fatalf("n=%d m=%d k=%d: expect doubled area %d, got %d", n, m, k, 2*n*m/k, area2(res.pts))
				}
			}
		}
	}
}
```

- [ ] **Step 3: Run the characterization tests against the current AC solver**

Run:

```bash
env GOCACHE=/tmp/learn-go-gocache go test -count=1 ./src/codeforces/set1/set10/set103/set1030/d/
```

Expected: `ok`, confirming that the test captures existing accepted behavior before refactoring.

### Task 2: Replace divisor enumeration with two GCD reductions

**Files:**
- Modify: `src/codeforces/set1/set10/set103/set1030/d/solution.go`
- Test: `src/codeforces/set1/set10/set103/set1030/d/solution_test.go`

- [ ] **Step 1: Replace `solve` with the standard construction**

Use this implementation while retaining the existing `result` type and `gcd` helper:

```go
func solve(n, m, k int) result {
	a, b := n, m

	g := gcd(a, k)
	a /= g
	k /= g

	g = gcd(b, k)
	b /= g
	k /= g

	if k > 2 {
		return result{ok: false}
	}

	if k == 1 {
		if 2*a <= n {
			a *= 2
		} else {
			b *= 2
		}
	}

	pts := [][]int{{0, 0}, {a, 0}, {0, b}}
	return result{ok: true, pts: pts}
}
```

- [ ] **Step 2: Format the Go files**

Run:

```bash
gofmt -w src/codeforces/set1/set10/set103/set1030/d/solution.go src/codeforces/set1/set10/set103/set1030/d/solution_test.go
```

Expected: both files remain valid `gofmt`-formatted Go.

- [ ] **Step 3: Run all package tests after the refactor**

Run:

```bash
env GOCACHE=/tmp/learn-go-gocache go test -count=1 ./src/codeforces/set1/set10/set103/set1030/d/
```

Expected: `ok`, including all samples and `TestSmall`.

### Task 3: Document and verify the final solution

**Files:**
- Modify: `src/codeforces/set1/set10/set103/set1030/d/problem.md`

- [ ] **Step 1: Replace `### ideas` with a complete solution explanation**

The new section must state:

- doubled area requires `a*b = 2*n*m/k`;
- `a=n` and `b=m` are reduced by `gcd(a,k)` and `gcd(b,k)`;
- if the remaining `k` is greater than two, construction is impossible;
- if it is two, the reduced sides are final;
- if it is one, double a reduced side that remains within its original bound;
- correctness follows because the reductions transfer every factor of the denominator into the two side lengths, and `k>=2` guarantees a side can be doubled in the remaining-one case;
- complexity is `O(log(max(n,m,k)))` time and `O(1)` extra space.

- [ ] **Step 2: Run final verification**

Run:

```bash
env GOCACHE=/tmp/learn-go-gocache go test -count=1 ./src/codeforces/set1/set10/set103/set1030/d/
git diff --check -- src/codeforces/set1/set10/set103/set1030/d/solution.go src/codeforces/set1/set10/set103/set1030/d/solution_test.go src/codeforces/set1/set10/set103/set1030/d/problem.md
```

Expected: the package reports `ok`, and `git diff --check` prints no whitespace errors.

- [ ] **Step 3: Review the final diff without committing user-owned work**

Run:

```bash
git diff -- src/codeforces/set1/set10/set103/set1030/d/solution.go src/codeforces/set1/set10/set103/set1030/d/solution_test.go src/codeforces/set1/set10/set103/set1030/d/problem.md
```

Expected: only the approved solver simplification, test improvements, and documentation update are present. Leave the package changes uncommitted for the user to review.
