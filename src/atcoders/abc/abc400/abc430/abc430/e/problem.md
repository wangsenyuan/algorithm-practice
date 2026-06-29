# E - Shift String

[Problem link](https://atcoder.jp/contests/abc430/tasks/abc430_e)

**Contest:** [AtCoder Beginner Contest 430](https://atcoder.jp/contests/abc430)

time limit: 2 sec

memory limit: 1024 MiB

score: 450 points

You are given strings `A` and `B` of equal length consisting of `0` and `1`.

You can perform the following operation on `A` zero or more times:

- Move the first character of `A` to the end.

Find the minimum number of operations required to make `A = B`. If it is impossible to make `A = B` no
matter how you operate, print `-1` instead.

You are given `T` test cases; find the answer for each of them.

## Constraints

- `1 <= T <= 10000`
- `A` and `B` are strings consisting of `0` and `1`.
- `2 <= |A| = |B| <= 10^6`
- For a single input, the sum of `|A|` does not exceed `10^6`.

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case is given in the following format:

```text
A
B
```

## Output

Print `T` lines. The `i`-th line should contain the answer for the `i`-th test case.

## Sample Input 1

```text
5
1010001
1000110
000
111
01010
01010
0101
0011
100001101110000001010110110001
101100011000011011100000010101
```

## Sample Output 1

```text
2
-1
0
-1
22
```

### Note

- In the first test case, `A = 1010001` and `B = 1000110`. After two operations,
  `1010001 -> 0100011 -> 1000110 = B`.
- In the second test case, `000` cannot become `111` by cyclic shifts.
- In the third test case, `A = B` from the start, so the answer is `0`.

### ideas
1. A + A 中要包括B
