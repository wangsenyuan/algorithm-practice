# GitHub Presentation Design

## Goal

Make the repository attractive and useful on GitHub without changing the competitive-programming solution tree. The first pass should present the repo as a portfolio-quality algorithm knowledge base: easy to understand from the root page, easy to browse by platform, and backed by CI that reports meaningful results.

## Scope

This design covers four public-facing surfaces:

- Root `README.md`
- Generated Markdown indexes under `docs/index/`
- A small stats/index generation tool
- The existing GitHub Actions workflow

The first implementation should not refactor solution packages, rename problem directories, introduce external Go dependencies, or build a static website.

## README

The root README should replace the generated template with a concise landing page:

- Short description of the repo and its purpose
- Current solution counts by platform
- Documentation coverage count for local `problem.md`, `readme.md`, and `README.md` files
- Directory layout with examples of common paths
- Commands for running a focused package test and a single sample test
- Links to generated platform indexes
- A short section explaining the local convention: each package is self-contained and many packages include a problem statement or explanation beside the solution

The README should avoid decorative inline SVGs and generic "show your support" boilerplate. Badges are useful only when they represent current repo state, such as license and CI.

## Indexes

Generated indexes should live under `docs/index/` and use plain Markdown so they render cleanly on GitHub.

The first pass should include:

- `docs/index/README.md`: index entry point and generation note
- `docs/index/codeforces.md`
- `docs/index/leetcode.md`
- `docs/index/codechef.md`
- `docs/index/atcoder.md`

Each platform index should list package paths and link to available local docs when present. It is acceptable for the first pass to be path-oriented rather than metadata-rich because many packages do not have structured titles or tags.

## Stats And Generation

Add one script that scans `src/` and regenerates both counts and indexes. It should use only the standard library or portable shell so the repo remains dependency-free.

The generator should report at least:

- Total `solution.go` files
- Total local docs files named `problem.md`, `readme.md`, or `README.md`
- Solution counts for Codeforces, LeetCode, CodeChef, and AtCoder

The generator may update a small generated stats file consumed by README maintenance. It does not need to edit `README.md` automatically in the first pass unless that keeps the workflow simpler.

## CI

The GitHub Actions workflow should stay focused on changed packages because the full repository test suite is too large for routine pull requests.

The workflow should:

- Keep testing package directories that contain changed Go files
- Fail the job when a changed package test fails
- Avoid swallowing failures with `|| true` on the actual `go test` command
- Keep enough logging to understand which packages were tested
- Avoid noisy branch-debug output unless it is needed for correctness

## Verification

Implementation is complete when:

- The generator can be run locally and produces deterministic Markdown indexes
- README links point to existing files
- Focused package testing still works for at least one existing package
- The CI workflow syntax remains valid YAML
- `git diff` shows no solution-code churn

## Risks And Non-Goals

Some indexes will initially be path-heavy because titles and topics are not consistently stored across all packages. That is acceptable for this first pass. Rich metadata, topic tags, difficulty labels, and a static website can come later after the basic public surface is useful.
