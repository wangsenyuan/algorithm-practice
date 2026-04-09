Given two strings `s` and `t` of length `n`, you aim to transform `s` into `t` through a series of the following operations:

- Construct a new string `s'` of length `n`, where `s'1 = s1`.
- For each `1 < i <= n`, `s'i` can be either `si` or `s(i-1)`.
- Then replace `s` with `s'`.

Your task is to achieve this transformation using the minimum number of operations. You also need to output the solution by printing the constructed string `s'` after each operation. If the transformation cannot be achieved in less than or equal to `kmax` operations, output `-1`.

## Input

Each test contains multiple test cases. The first line contains the number of test cases `t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains two integers `n`, `kmax` (`1 <= n * kmax <= 10^6`) — the length of two strings and the maximum number of operations that can be used.

The second line of each test case contains one string `s` of length `n`.

The third line of each test case contains one string `t` of length `n`.

It is guaranteed that the sum of `n * kmax` over all test cases does not exceed `10^6`.

It is guaranteed that both `s` and `t` consist of lowercase Latin letters.

## Output

For each test case:

- If you cannot achieve the transformation in less than or equal to `kmax` operations, simply output `-1` in one line.
- Otherwise, on the first line, output one integer `k <= kmax` — the minimum number of operations. Then `k` lines follow, each line contains one string of length `n` — the string after each operation.
- If there are multiple solutions, output any.

## Example

### Input

```text
7
4 1
abcd
aabd
2 2
ab
ab
5 3
abcde
abbcc
9 1
egcnyeluw
eegccyelw
10 3
vzvylxxmsy
vvvvvllxxx
4 6
acba
aaac
5 7
acabb
aaaca
```

### Output

```text
1
aabd
0
2
abbcd
abbcc
-1
3
vvzvylxxms
vvvzvllxxm
vvvvvllxxx
2
aacb
aaac
2
aacab
aaaca
```

## Note

In the first test case, obviously `s` can be transformed to `t` in one operation.

In the second test case, initially `s = t`, so no operation is needed.

In the fourth test case, although `s` can be transformed to `t` in two operations, `kmax = 1`, so the answer is `-1`.

### ideas
1. s[1]不能被改变（所以 s[1] != t[1] => -1)
2. 考虑s[n], t[n], 如果s[n] != t[n], 那么必须找到一个s[i] = t[n], 通过n-i次操作将s[n]编程s[i]
3. 如果n-i > kmx => -1
4. 同时 i到n中间的都需要s[i]

### solution summary

Scan positions from right to left and greedily fix `t[i]`.

Key observations:

1. `s[0]` never changes, so if `s[0] != t[0]`, answer is impossible.
2. In one operation, each position can only copy from itself or left neighbor, so characters only propagate to the right.
3. To make `i` become `t[i]`, we must pick some `i1 <= i` with current `s[i1] = t[i]`, then shift that value right step by step; this needs exactly `i - i1` operations.

Greedy construction:

- Maintain `buf` as current string state during planning.
- Maintain `pos[c]`: indices where character `c` currently appears in `buf` (as stacks, rightmost accessible index at the end).
- For `i = n-1 .. 1`:
  - Remove index `i` from `pos[buf[i]]` (this position is being finalized).
  - If `buf[i] == t[i]`, continue.
  - Otherwise choose `i1 = rightmost index in pos[t[i]]`.
    - If not found, impossible.
    - If `i - i1 > kmax`, impossible.
  - Record that for every distance `d = 1..i-i1`, position `i1+d` must copy from left in operation `d`.
  - Simulate these shifts on `buf` and update `pos` so later positions use the latest state.

Let `k` be the maximum needed distance over all fixed positions. This `k` is minimal, because each fixed position requiring distance `d` enforces at least `d` operations globally.

Finally, replay operations `1..k` on the original `s` using recorded position lists to output each intermediate string.

Complexity:

- Each index is pushed/popped/updated `O(1)` amortized in the right-to-left process.
- Total time `O(n)`, memory `O(n)` per test case.