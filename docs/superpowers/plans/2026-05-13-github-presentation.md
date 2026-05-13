# GitHub Presentation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a polished GitHub-facing documentation layer for the competitive-programming repository.

**Architecture:** Keep all solution packages unchanged. Add a small standard-library Go generator under `tools/indexgen/` that scans `src/`, computes stats, and writes Markdown indexes under `docs/index/`; then rewrite the root README to link to those generated files and clean up the existing changed-package CI workflow.

**Tech Stack:** Go standard library, Markdown, GitHub Actions YAML.

---

## File Map

- Create `tools/indexgen/main.go`: CLI entry point and generator implementation.
- Create `tools/indexgen/main_test.go`: fixture-based tests for path scanning, stats, docs detection, and deterministic Markdown output.
- Create `docs/index/README.md`: generated index landing page.
- Create `docs/index/codeforces.md`: generated Codeforces package index.
- Create `docs/index/leetcode.md`: generated LeetCode package index.
- Create `docs/index/codechef.md`: generated CodeChef package index.
- Create `docs/index/atcoder.md`: generated AtCoder package index.
- Modify `README.md`: polished GitHub landing page.
- Modify `.github/workflows/go.yml`: remove noisy debug output and fail on changed-package test failures.

## Task 1: Generator Tests

**Files:**
- Create: `tools/indexgen/main_test.go`

- [x] **Step 1: Write fixture-based tests before implementation**

Create tests that build a temporary fake repo with these files:

```text
src/codeforces/set1/set18/set185/set1857/g/solution.go
src/codeforces/set1/set18/set185/set1857/g/problem.md
src/leetcode/set1000/set3000/set3900/set3920/p3921/solution.go
src/codechef/easy/section00/example/solution.go
src/codechef/easy/section00/example/readme.md
src/atcoders/arc/arc100/arc127/d/solution.go
src/unknown/demo/solution.go
```

Assert:

- total solutions is `5`
- total docs is `2`
- platform counts are Codeforces `1`, LeetCode `1`, CodeChef `1`, AtCoder `1`
- unknown solutions are counted in total but not platform index totals
- generated Codeforces Markdown links to the package and its `problem.md`

- [x] **Step 2: Run tests and confirm RED**

Run: `go test ./tools/indexgen`

Expected: FAIL because `collectRepo` and Markdown rendering functions do not exist yet.

## Task 2: Generator Implementation

**Files:**
- Create: `tools/indexgen/main.go`
- Modify: `tools/indexgen/main_test.go`

- [x] **Step 1: Implement minimal generator**

Implement:

- `collectRepo(root string) (RepoIndex, error)`
- `renderIndexHome(index RepoIndex) string`
- `renderPlatform(index RepoIndex, platform Platform) string`
- `writeIndexes(root string, index RepoIndex) error`
- `main()`

Use `filepath.WalkDir`, normalize output paths with `filepath.ToSlash`, sort entries lexicographically, and write files with mode `0644`.

- [x] **Step 2: Run generator tests and confirm GREEN**

Run: `go test ./tools/indexgen`

Expected: PASS.

- [x] **Step 3: Generate repository indexes**

Run: `go run ./tools/indexgen`

Expected generated files:

```text
docs/index/README.md
docs/index/codeforces.md
docs/index/leetcode.md
docs/index/codechef.md
docs/index/atcoder.md
```

## Task 3: README Rewrite

**Files:**
- Modify: `README.md`

- [x] **Step 1: Replace generated template**

Write a concise README with:

- project purpose
- CI and license badges
- stats table using the generated counts
- links to generated indexes
- directory layout examples
- focused test commands
- self-contained package convention

- [x] **Step 2: Verify README links locally**

Run: `rg -n "\\]\\(([^)#]+)\\)" README.md docs/index/README.md`

Expected: links point to files that exist in the repo.

## Task 4: CI Cleanup

**Files:**
- Modify: `.github/workflows/go.yml`

- [x] **Step 1: Simplify changed-file detection and test execution**

Keep pull-request scope. Use `git diff --name-only origin/${{ github.base_ref }}...HEAD` after fetching the base branch. Build a sorted unique list of directories containing changed `.go` files.

- [x] **Step 2: Remove swallowed test failures**

The package test command must not use `|| true`. If any changed package test fails, the job should fail.

- [x] **Step 3: Validate YAML shape**

Run: `go test ./tools/indexgen`

Expected: PASS. Also inspect `.github/workflows/go.yml` for valid indentation and no `go test ... || true`.

## Task 5: Final Verification

**Files:**
- All changed files

- [x] **Step 1: Re-run generator**

Run: `go run ./tools/indexgen`

Expected: no unexpected changes after generation.

- [x] **Step 2: Run focused generator tests**

Run: `go test ./tools/indexgen`

Expected: PASS.

- [x] **Step 3: Run one existing package test**

Run: `go test ./src/codeforces/set1/set18/set185/set1857/g/`

Expected: PASS.

- [x] **Step 4: Confirm solution packages were untouched**

Run: `git diff --name-only`

Expected: only README, docs, tool, and workflow files changed.
