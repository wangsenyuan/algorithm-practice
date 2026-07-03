# G. Inverse Function

[Problem link](https://codeforces.com/problemset/problem/39/G)

**Contest:** [Codeforces Beta Round #10 (Div. 2)](https://codeforces.com/contest/39)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Petya wrote a C++ program that computes an interesting function `f(n)`. When he returned, the program had
already printed `f(n)`, but the input value `n` was lost.

Given the output value and the source code of `f`, find the maximum integer `n` in `[0, 32767]` such
that running the program on `n` produces the given output. If no such `n` exists, print `-1`.

The function description may contain extra spaces and line breaks, but the keywords `int`, `if`, `return`
and numeric literals remain intact. The total input size does not exceed 100 bytes.

## Input

The first line contains the printed value `f(n)` (`0 <= f(n) <= 32767`), immediately followed by the
start of the function source code on the same line. Any remaining source code continues on following
lines until EOF.

## Output

Print one integer — the maximum valid input `n`, or `-1` if none exists.

## Example

### Input

```text
17int f(int n){if (n < 100) return 17;if (n > 99) return 27;}
```

### Output

```text
99
```

### Input

```text
13int f(int n){if (n == 0) return 0;return f(n - 1) + 1;}
```

### Output

```text
13
```

### Input

```text
144int f(int n){if (n == 0) return 0;if (n == 1) return n;return f(n - 1) + f(n - 2);}
```

### Output

```text
24588
```

## Solution

The input is a tiny program, so the direct approach is to parse and interpret the restricted language.

First remove all whitespace from the function source. Whitespace is optional in the grammar and cannot
split keywords or numbers, so this does not change the program. Then parse the body between `{` and
`}` as a sequence of operators:

1. `return arithmExpr;`
2. `if (logicalExpr) return arithmExpr;`

Arithmetic expressions are parsed with recursive descent:

1. `parseSum` handles `+` and `-`.
2. `parseProduct` handles `*` and `/`.
3. `parseMultiplier` handles `n`, constants, and recursive calls `f(arithmExpr)`.

The parser builds a small syntax tree for each expression and condition. The evaluator executes the
operators in order. For an `if`, it evaluates the condition and only returns the expression when the
condition is true. For a plain `return`, it immediately returns the expression value.

The arithmetic follows the 15-bit compiler rules:

1. Addition, subtraction, and multiplication are computed modulo `32768`.
2. Subtraction is normalized back into `[0, 32767]`.
3. Division is ordinary integer division.

For recursion, store `f(n)` in a memo array of length `32768`. The statement guarantees that calls
from `f(N)` only use smaller arguments, so evaluating values from `0` to `32767` is safe and each
`f(n)` is computed once. Whenever `f(n)` equals the target value, update the answer to `n`; because
the scan is increasing, the final answer is the maximum valid `n`.

### Correctness

The parser follows the exact grammar of the restricted C++ function: it preserves operator precedence
by parsing products inside sums, and it parses recursive calls as multipliers with their own arithmetic
argument. Therefore each expression tree represents the same expression as the source program.

The evaluator processes the parsed operators in source order. A conditional return is used exactly
when its logical expression is true, and a plain return always stops execution. This is the same
control flow as the function definition, so for any fixed `n`, the evaluator returns the same value as
the given program.

Each arithmetic operation is evaluated with the required 15-bit rules, and recursive calls are resolved
through the same evaluator. Since the statement guarantees that recursive calls only use smaller
arguments, memoization does not change the result; it only avoids recomputing already determined
values.

The algorithm checks every possible input `n` in `[0, 32767]`. It records exactly those `n` whose
computed value equals the target and keeps the largest one. Thus it returns the required maximum valid
input, or `-1` if no checked value matches.

### Complexity

The source size is at most `100` bytes, and there are only `32768` possible input values. Each `f(n)`
is memoized after its first evaluation, so the running time is `O(32768 * L)`, where `L` is the parsed
program size. The memory usage is `O(32768 + L)`.
