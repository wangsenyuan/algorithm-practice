# Algorithm Practice In Go

[![Go CI/CD](https://github.com/wangsenyuan/algorithm-practice/actions/workflows/go.yml/badge.svg)](https://github.com/wangsenyuan/algorithm-practice/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)

Competitive-programming solutions in Go, organized as self-contained packages by online judge and contest. The repository is mostly standard-library Go: each problem directory owns its parser, `solve` function, tests, and, when available, a local statement or explanation file.

## Repository Stats

Generated with `go run ./tools/indexgen`.

| Area | Count |
| --- | ---: |
| Total solutions | 10256 |
| Local problem docs and explanations | 3075 |
| Codeforces solutions | 4655 |
| LeetCode solutions | 3016 |
| CodeChef solutions | 2252 |
| AtCoder solutions | 155 |

## Browse

- [All generated indexes](docs/index/README.md)
- [Codeforces](docs/index/codeforces.md)
- [LeetCode](docs/index/leetcode.md)
- [CodeChef](docs/index/codechef.md)
- [AtCoder](docs/index/atcoder.md)

The generated indexes are path-oriented on purpose. Many older packages do not have normalized titles or topic tags, but the package paths are stable and link directly to local `problem.md`, `readme.md`, or `README.md` files when those docs exist.

## Layout

```text
src/
  codeforces/set1/set18/set185/set1857/g/
    solution.go
    solution_test.go
  leetcode/set1000/set3000/set3900/set3920/p3921/
    solution.go
    solution_test.go
  codechef/easy/section00/example/
    solution.go
    solution_test.go
  atcoders/arc/arc100/arc127/d/
    solution.go
    solution_test.go
```

Most packages use `package main`. `solution.go` contains the production stdin/stdout path and the algorithm; `solution_test.go` calls the same parsing path or `solve` function with sample data.

## Run Tests

Run one problem package:

```bash
go test ./src/codeforces/set1/set18/set185/set1857/g/
```

Run one sample test inside a package:

```bash
go test ./src/codeforces/set1/set18/set185/set1857/g/ -run TestSample1
```

Regenerate the public indexes:

```bash
go run ./tools/indexgen
```

Running `go test ./...` is intentionally not the default workflow; this repository contains thousands of packages, so focused package tests are the normal development path.

## Conventions

- Solution packages are self-contained and do not import shared helper packages.
- I/O helpers are copied into each package when needed.
- Local statement or explanation files live beside the solution as `problem.md`, `readme.md`, or `README.md`.
- New documentation should describe the accepted reasoning in the same notation used by the local solution.

## License

Apache License 2.0. See [LICENSE](LICENSE).
