# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

A Go competitive programming practice repository solving problems from LeetCode, Codeforces, CodeChef, AtCoder, CodeJam, CodeWars, Luogu, and Gym contests. No external dependencies — pure Go standard library.

## Commands

```bash
# Run all tests in a specific problem's package
go test ./src/codeforces/set1/set18/set185/set1857/g/

# Run a single test case
go test ./src/codeforces/set1/set18/set185/set1857/g/ -run TestSample1

# Run all tests (slow — tens of thousands of packages) must ask permission before execution
go test ./...

# Update Go version
go mod edit -go 1.21
go mod tidy
```

## Code Structure

Each problem lives in its own directory as `package main`:
- `solution.go` — `main()` reads from stdin/writes to stdout; `solve(...)` contains the algorithm
- `solution_test.go` — calls `solve(...)` directly with expected outputs

### Directory Conventions

- **Codeforces**: `src/codeforces/set{0,1,2}/set{XY}/set{XY0}/set{XYZW}/` where `XYZW` is the contest number, then `{a,b,c,...}/` for problem letter
- **LeetCode**: `src/leetcode/set1000/set{N000}/set{N0X0}/p{NXYZ}/`
- **CodeChef**: `src/codechef/{easy,medium}/section{NN}/...`
- **AtCoder**: `src/atcoders/arc/arc{N00}/arc{NXY}/`
- **Gym**: `src/gym/set{contest_id}/`

### Solution Template Pattern

Every `solution.go` embeds its own I/O helpers (no shared utility package). The preferred pattern separates I/O into a `drive()` function:

```go
package main

import ("bufio"; "fmt"; "os")

func main() {
    reader := bufio.NewReader(os.Stdin)
    res := drive(reader)
    fmt.Println(res)
}

func drive(reader *bufio.Reader) <type> {
    // parse input using fmt.Fscan or custom readX helpers
    return solve(...)
}

// I/O helpers (when fmt.Fscan is insufficient): readNum, readTwoNums, readNNums, readInt, readInt64, etc.
// Math helpers (when needed): mul, add, pow (modular arithmetic, typically mod=998244353 or 1e9+7)
// Data structures (when needed): UFSet (union-find), segment trees, etc.

func solve(...) <type> {
    // algorithm
}
```

### Test Pattern

Tests pass raw input strings to `drive()` — the same I/O path used in production:

```go
package main

import (
    "bufio"
    "reflect"
    "strings"
    "testing"
)

func runSample(t *testing.T, s string, expect <type>) {
    reader := bufio.NewReader(strings.NewReader(s))
    res := drive(reader)
    if !reflect.DeepEqual(res, expect) {
        t.Errorf("Sample expect %v, but got %v", expect, res)
    }
}

func TestSample1(t *testing.T) {
    runSample(t, "input line here\n", expectedValue)
}
```

### Notes

- Each `solution.go` is self-contained — copy I/O helpers into every new file rather than importing a shared package.
- Problem statements are sometimes stored as `problem.md` or `readme.md` alongside the solution.
- Some solutions include Chinese comments explaining the algorithm approach.
