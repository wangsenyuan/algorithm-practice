# ABC265 C Scaffold Design

## Scope

Add a new package at `src/atcoders/abc/abc200/abc260/abc265/c` for AtCoder ABC265 C, "Belt Conveyor". The package will contain the English problem statement, production I/O scaffolding, and the three official sample tests. The algorithm itself is intentionally excluded.

## Files

- `problem.md`: English statement, constraints, input/output format, and official samples.
- `solution.go`: a self-contained `package main` with `main`, `drive`, input parsing, output formatting, and a `solve(h, w int, grid []string) []int` function marked `TODO`.
- `solution_test.go`: exercises the production `drive` path with all three official samples. Tests call `t.Skip("solve TODO")` until the algorithm is implemented, so the unfinished scaffold still has a clean package test result.

## Output Contract

`solve` returns `{row, column}` for a finite endpoint and `{-1}` for an infinite path. `drive` formats either result as required by the problem.

## Verification

Run `gofmt` on Go files and the package-local `go test`. Confirm that only the intended new package and this design document are changed.
