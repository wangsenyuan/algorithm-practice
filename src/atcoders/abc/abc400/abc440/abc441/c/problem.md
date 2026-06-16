# C - Sake or Water (ABC441)

**Contest:** [ABC441](https://atcoder.jp/contests/abc441) — AtCoder Beginner Contest 441  
**Task:** [https://atcoder.jp/contests/abc441/tasks/abc441_c](https://atcoder.jp/contests/abc441/tasks/abc441_c)

**Time limit:** 2 sec / **Memory limit:** 1024 MiB  
**Score:** 300 points

## Problem Statement

There are `N` cups. Cup `i` contains `A_i` ml of colorless liquid.

Exactly `K` cups contain sake; the other `N - K` cups contain water. Which cups
hold sake is unknown.

Takahashi chooses one or more cups and drinks all liquid in them. Find the
**minimum number of cups** he must choose so that, no matter which `K` cups contain
sake, he drinks at least `X` ml of sake.

If no such choice exists, print `-1`.

## Constraints

- `1 <= K <= N <= 3 * 10^5`
- `1 <= A_i <= 10^9`
- `1 <= X <= 3 * 10^14`
- All input values are integers

## Input

The input is given from Standard Input in the following format:

```text
N K X
A_1 A_2 ... A_N
```

## Output

Print the minimum number of cups needed. If impossible, print `-1`.

## Sample Input 1

```text
3 2 5
10 6 8
```

## Sample Output 1

```text
2
```

If Takahashi chooses cups 1 and 3:

- Sake in cups 1 and 2: he drinks `10` ml sake and `8` ml water
- Sake in cups 1 and 3: he drinks `18` ml sake
- Sake in cups 2 and 3: he drinks `8` ml sake and `10` ml water

In every case he gets at least `5` ml of sake. Choosing only one cup cannot
guarantee that, so the answer is `2`.

## Sample Input 2

```text
2 1 8
6 10
```

## Sample Output 2

```text
-1
```

If cup 1 contains sake, no choice of cups guarantees at least `8` ml of sake.

## Sample Input 3

```text
5 3 3000000000
1000000000 1000000000 1000000000 1000000000 1000000000
```

## Sample Output 3

```text
5
```

### ideas
1. 在喝完w个杯子后,保证至少喝到了Xml的sake.
2. 排序后,最前面的v个不是的情况下,中间的k-v个杯子 >= X