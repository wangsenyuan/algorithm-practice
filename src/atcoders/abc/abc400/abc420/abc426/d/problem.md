# D - Pop and Insert

[Problem link](https://atcoder.jp/contests/abc426/tasks/abc426_d)

**Contest:** [AtCoder Beginner Contest 426](https://atcoder.jp/contests/abc426)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

## Problem

You are given a string `S` of length `N` consisting of `0` and `1`.

You may perform the following operation any number of times (including zero):

- Delete the first or last character, flip it (`0` ↔ `1`), and insert it back at
  any position. Formally, with `r(0)=1` and `r(1)=0`:
  - Choose `i` (`1 <= i <= N`) and change `S` to
    `S_2 ... S_i r(S_1) S_{i+1} ... S_N`.
  - Choose `i` (`0 <= i <= N-1`) and change `S` to
    `S_1 ... S_i r(S_N) S_{i+1} ... S_{N-1}`.

Find the minimum number of operations to make all characters of `S` the same.
Such a sequence of operations always exists.

There are `T` test cases; solve each of them.

## Constraints

- `1 <= T <= 2 * 10^5`
- `2 <= N <= 5 * 10^5`
- `T` and `N` are integers
- `S` is a string of length `N` consisting of `0` and `1`
- The sum of `N` over all test cases is at most `5 * 10^5`

## Input

```text
T
case_1
case_2
...
case_T
```

Each test case:

```text
N
S
```

## Output

Print `T` lines. The `i`-th line is the answer for the `i`-th test case.

## Sample 1

```text
Input
3
5
01001
3
000
15
110010111100101

Output
4
0
16
```


### ideas
1. 假设最终的字符串是0
2. 如果中间有1, 那么就必须从左端(或者右端)移动到这个位置, 把它变成0
3. 假设最后的结果是以s[i]为目标(和中心)的, 那么它左边的不等于s[i]的, 需要一次操作
4. 而和它一样的, 需要两次操作(而且只需要考虑最近的那个)

## Solution Summary

### Choose the untouched middle run

Think of every character as a distinct object. An operation can take an object
only when it is at the left or right end.

In an optimal process, leave a non-empty middle part untouched and move the
characters outside it. Because characters can be removed only from the two
ends, the untouched characters form a contiguous substring of the original
string. Since the final string is constant, every character in this untouched
substring must already have the same value. Therefore, it can be extended to
one maximal run of equal characters without increasing the answer.

Suppose the chosen run consists of the bit `d`.

For every character outside this run:

- if its value is `1-d`, one operation is enough: remove it from an end, flip
  it to `d`, and insert it into its final position;
- if its value is `d`, it cannot be left untouched because it lies outside the
  chosen middle run. Each operation flips it, so it needs at least two
  operations to finish as `d`.

These costs are also achievable by processing characters from both ends toward
the chosen run. Thus, for a fixed run of `d`, the cost is:

```text
(number of 1-d outside the run)
+ 2 * (number of d outside the run)
```

It remains to evaluate this expression for every maximal run.

### Suffix costs

For each position `i`, define `suf[i]` as the cost of processing all characters
strictly to the right of `i`, using `s[i]` as the target bit:

```text
suf[i] =
    count of 1-s[i] in S[i+1..N)
    + 2 * count of s[i] in S[i+1..N)
```

The code builds this array from right to left. Before processing `i`, `cnt[0]`
and `cnt[1]` contain the character counts in the suffix:

```go
d := int(s[i] - '0')
suf[i] = cnt[1-d] + 2*cnt[d]
cnt[d]++
```

### Enumerating maximal runs

Scan the string from left to right by maximal runs. Let the current run be
`S[j..i)` and let its bit be `d`.

At this moment, `cnt` contains the counts strictly to the left of the run, so
the left-side cost is:

```text
cnt[1-d] + 2 * cnt[d]
```

The last character of the run is at `i-1`. Since `suf[i-1]` excludes that
character and counts everything to its right using the same target `d`, it is
exactly the right-side cost.

Therefore, the candidate answer for this run is:

```text
cnt[1-d] + 2*cnt[d] + suf[i-1]
```

Take the minimum candidate over all runs, then add the current run's length to
the corresponding prefix count before examining the next run.

### Correctness

For any chosen target run of bit `d`:

1. The run itself needs no operations.
2. Every outside `1-d` character needs an odd number of flips, so at least one
   operation; one operation is sufficient.
3. Every outside `d` character must be moved but must finish unchanged, so it
   needs a positive even number of flips, hence at least two operations; two
   operations are sufficient.

The computed candidate is therefore both a lower bound and achievable for that
run. Every possible untouched middle part belongs to some maximal run, and the
algorithm checks every maximal run. Hence the minimum candidate is the global
minimum number of operations.

### Complexity

The suffix construction and run scan are both linear:

```text
Time:  O(N) per test case
Space: O(N)
```

Because the sum of `N` over all test cases is at most `5 * 10^5`, the total
time is linear in the input size.
