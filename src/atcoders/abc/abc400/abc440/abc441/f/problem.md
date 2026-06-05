# F - Must Buy (ABC441)

**Contest:** [ABC441](https://atcoder.jp/contests/abc441) — AtCoder Beginner Contest 441  
**Task:** [https://atcoder.jp/contests/abc441/tasks/abc441_f](https://atcoder.jp/contests/abc441/tasks/abc441_f)

**Time limit:** 2.5 sec / **Memory limit:** 1024 MiB  
**Score:** 500 points

## Problem Statement

A store sells `N` items, numbered item `1`, item `2`, ..., item `N`. Item `i`
(`1 <= i <= N`) has price `P_i` yen and value `V_i`. Each item appears only once.

Takahashi chooses some items so that the total price is at most `M` yen. It is
guaranteed that `M` is at least the price of every item. That is, for every
`1 <= i <= N`, there exists a valid selection with total price at most `M` that
includes item `i`.

For each `1 <= i <= N`, classify item `i` into one of three categories:

- **Category A:** To maximize the total value of chosen items (with total price at
  most `M`), item `i` must be chosen.
- **Category B:** To maximize the total value, item `i` may be chosen or not chosen.
- **Category C:** To maximize the total value, item `i` must never be chosen.

## Constraints

- `1 <= N <= 1000`
- `1 <= M <= 5 * 10^4`
- `1 <= P_i <= M`
- `1 <= V_i <= 10^9`
- All input values are integers.

## Input

The input is given from standard input in the following format:

```text
N M
P_1 V_1
P_2 V_2
:
P_N V_N
```

## Output

Print a string of length `N` consisting of `A`, `B`, and `C`. If item `i` is in
category `X` (`X` is `A`, `B`, or `C`), the `i`-th character of the output must
be `X`.

## Sample Input 1

```text
5 7
2 5
2 5
3 5
3 10
3 20
```

## Sample Output 1

```text
BBCBA
```

The maximum total value with total price at most `7` is `30`, achieved only by:

- choosing items `1`, `2`, and `5`, or
- choosing items `4` and `5`.

Thus:

- item `5` is category **A** (must choose),
- items `1`, `2`, and `4` are category **B** (optional),
- item `3` is category **C** (never choose).

So the answer is `BBCBA`.

## Sample Input 2

```text
7 3
1 1
1 1
1 2
1 2
1 2
1 3
1 3
```

## Sample Output 2

```text
CCBBBAA
```

The maximum total value is achieved by choosing any one of items `3`, `4`, or `5`,
together with items `6` and `7`. Even if two items have the same price and value,
they are still distinct items.
