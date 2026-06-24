# K. Kindergarten Homework

[Problem link](https://codeforces.com/problemset/problem/2172/K)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Kenny is a kindergarten teacher who is teaching his students about math. To help them
learn, he wants to give them homework in an enjoyable and non-annoying format — Number
Search.

A Number Search puzzle is similar to a Word Search puzzle, but instead of words, it is
about finding math expressions. An instance of the puzzle consists of two parts: a grid
with `n` rows and `m` columns, and a list `a_1, a_2, ..., a_q` of `q` target numbers to
search for. Each cell in the grid contains a digit from `1` to `9`, a plus sign `+`, or a
multiplication sign `*`.

An expression can be formed by joining one or more consecutive characters along the same
row, column, or diagonal, in a straight line (in either direction). Each expression can be
defined by its starting cell and ending cell. Two expressions are considered different if
they start or end at different cells, even if they contain the exact same characters.

A valid expression must be of the form `num op num op ... op num`, where each `num` is a
sequence of one or more digits, and each `op` is either `+` or `*`. For example, `123` and
`1+2*3+45` are valid expressions, while `+123` and `2**5` are not. The value of an
expression is the result of evaluating it following typical arithmetic rules (multiplication
before addition).

The goal of the puzzle is to count, for each number `a_i`, how many different valid
expressions in the grid evaluate to `a_i`.

## Input

The first line contains three integers `n`, `m`, and `q` — the number of rows, the number
of columns, and the number of target numbers.

Each of the following `n` lines contains a string of `m` characters, representing the grid.

The `i`-th of the following `q` lines contains an integer `a_i`, representing the `i`-th
target number.

## Constraints

- `1 <= n, m <= 1000`
- `1 <= n * m <= 3 * 10^4`
- `1 <= q <= 10^5`
- The grid contains only characters in `123456789+*`
- `1 <= a_i <= 10^6`
- All `a_i` are distinct

## Output

Print `q` lines. The `i`-th line contains an integer — the answer for `a_i`.

## Example

### Input

```text
9 8 4
4216793+
717*4*54
7+5+727*
45149+71
8+26697*
+189*2+9
5+*244+7
42595952
97+*315+
67
420
3
727
```

### Output

```text
2
0
4
3
```

## Note

Let `(r, c)` denote the cell at row `r` and column `c` (1-indexed).

The answer for `3` is `4`:

- The expression `3` appears twice, at `(1, 7)` and `(9, 5)`.
- The expression `1+2` appears from `(6, 2)` to `(8, 2)`.
- The expression `2+1` appears from `(8, 2)` to `(6, 2)`. Even though they share the same
  set of cells, the two expressions are different because they start at different positions.

The answer for `727` is `3`:

- The expression `7+16*45` appears from `(2, 1)` to `(8, 7)`. Multiplication is performed
  before addition.
- The expression `727` appears from `(3, 5)` to `(3, 7)`, and from `(3, 7)` to `(3, 5)`.
  Even though they share the same set of cells and characters, they are different because
  they start at different positions.

## Solution

Only target values up to `10^6` matter. Every digit is positive and both operators are
monotone for positive operands, so once a currently built number or reduced expression
already exceeds `10^6`, extending it cannot make it useful again. This allows us to
enumerate candidate expressions directly and stop early.

An expression can go in any of 8 straight directions. Instead of writing separate code
for every direction, process only two directions on the current grid:

- left to right on each row;
- down-right on each diagonal.

Then rotate the whole grid by 90 degrees and repeat this four times. Across the four
rotations, these two directions cover all rows, columns, and diagonals in both possible
orientations. Single-cell expressions are counted separately once, because rotations
would otherwise count the same cell multiple times.

For each starting cell on a processed line, scan forward one character at a time and
maintain the value of the expression prefix ending at the current cell:

- If the prefix starts with an operator, it is invalid.
- If two operators appear consecutively, the scan stops for this start position.
- Digits extend the current number, or start a new number immediately after an operator.
- Every time the current character is a digit and the prefix is valid, its value is
  counted if it is at most `10^6`.

The implementation uses a tiny expression stack to respect multiplication precedence.
Before pushing a new operator, it reduces the previous operation when the previous
operator has at least the same precedence:

```text
previous '*' before anything: reduce
previous '+' before new '+': reduce
previous '+' before new '*': keep it pending
```

Because only `+` and `*` exist, the stack never needs to hold many terms. It is enough
to keep the already reduced left part, the pending operator, and the current product
part. When a prefix ends at a digit, the code can compute its value from this small
stack and increment the frequency table for that value.

After all rotations have been processed, `dp[x]` stores the number of valid multi-cell
expressions with value `x`, and the separate digit table stores single-cell expressions.
Each query `a_i` is answered by adding these counts.

## Correctness

We prove that the algorithm outputs the number of valid expressions evaluating to each
target value.

First, every valid expression lies on a row, column, or diagonal, in one of two
directions along that line. In one of the four grid rotations, that directed line becomes
either a left-to-right row segment or a down-right diagonal segment. Therefore the
enumeration over the two directions in all four rotations visits every multi-cell
expression at least once. It visits each directed start/end pair exactly once, because a
specific direction becomes one of the two processed directions in exactly one rotation.
Single-cell expressions are intentionally excluded from the rotated enumeration and
counted once per digit cell, so they are also counted exactly once.

For a fixed start position and direction, the scan considers all prefixes ending at
later cells until no useful valid expression can be formed. A prefix beginning with an
operator or containing two consecutive operators is not a valid expression, and adding
more characters after two consecutive operators cannot make any prefix ending before the
second operator valid, so stopping there does not lose a valid expression for this start.
If a numeric value grows past `10^6`, every future extension uses only positive digits
and the monotone operators `+` and `*`, so the value cannot decrease into the query
range; stopping there also loses no answer.

The stack evaluation is the standard left-to-right evaluation with multiplication before
addition. Whenever a new operator is read, all pending operations with precedence at
least as high as the new one are reduced. Thus multiplication is reduced before addition,
and additions are reduced in left-to-right order. Consequently, every valid prefix ending
at a digit is counted under its true arithmetic value.

Combining the direction coverage, exact per-prefix validation, and correct expression
evaluation, the frequency table contains exactly the number of valid expressions for
each value. The algorithm returns those frequencies for the requested targets, so every
answer is correct.

## Complexity

For each rotation and each processed line, the algorithm starts a scan from every cell.
Each scan stops when it becomes invalid or when the value exceeds `10^6`; in the worst
case this is bounded by the line length. With `n * m <= 3 * 10^4` and line length at
most `max(n, m)`, the enumeration fits the required limits.

Time complexity: `O(n * m * max(n, m))` in the worst case.

Memory complexity: `O(10^6 + n * m)`.
