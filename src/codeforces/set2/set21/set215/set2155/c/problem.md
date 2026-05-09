C. The Ancient Wizards' Capes

Time limit per test: 2 seconds  
Memory limit per test: 256 megabytes

There are `n` wizards in a row numbered `1` to `n` from left to right. Each wizard has an invisibility cape that can be worn either on his left side or on his right side.

Harry walks from the position of wizard `1` to the position of wizard `n` (`1 <= n <= 10^5`) and records how many wizards he sees from each position.

A wizard at position `j` is visible from position `i` if:

- Wizard `j` wears his cape on his left side and `i >= j`.
- Wizard `j` wears his cape on his right side and `i <= j`.

In particular, wizard `i` is visible from position `i`.

Harry's list is an array `a` of `n` elements, where `a_i` (`1 <= a_i <= n`) is the number of wizards Harry saw from the position of wizard `i`.

Determine how many cape arrangements are consistent with this list, modulo `676767677`.

## Input

Each test contains multiple test cases.

- The first line contains the number of test cases `t` (`1 <= t <= 10^4`).

For each test case:

- The first line contains one integer `n` (`1 <= n <= 10^5`) — the length of `a`.
- The second line contains `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= n`).

It is guaranteed that the sum of `n` over all test cases does not exceed `10^5`.

## Output

For each test case, print one integer: the number of valid cape arrangements modulo `676767677`.

## Example

### Input

```text
7
1
1
4
4 4 3 2
3
1 3 2
2
2 1
3
2 2 3
3
2 3 3
3
2 2 2
```

### Output

```text
2
1
0
1
2
0
0
```

## Note

The image in the original statement shows one arrangement that matches the second test case:

- Wizard `1` wears the cape on the left side.
- Wizards `2`, `3`, and `4` wear the cape on the right side.

Then:

- From position `1`, we see wizards `1,2,3,4`.
- From position `2`, we see wizards `1,2,3,4`.
- From position `3`, we see wizards `1,3,4`.
- From position `4`, we see wizards `1,4`.

So Harry's list is `[4,4,3,2]`, and this arrangement is unique.

In the third test case, no arrangement can produce Harry's list.

In the fifth test case, there are two valid arrangements:

- `1||23|`
- `|12||3`

Original source: [Codeforces 2155C](https://codeforces.com/problemset/problem/2155/C)

## Solution Notes

Encode each wizard's cape direction as:

- `0`: the cape is on the left side, so this wizard is visible from positions on or after him.
- `1`: the cape is on the right side, so this wizard is visible from positions on or before him.

For a fixed position `i`, the number of visible wizards is:

```text
left-facing wizards before i + wizard i itself + right-facing wizards after i
```

The current wizard is always visible, regardless of his cape direction.

The important observation is that adjacent values almost determine adjacent cape directions. When moving from position `i` to position `i + 1`, all far-away wizards keep the same visibility. Only wizard `i` and wizard `i + 1` can change their contribution:

- Wizard `i` remains visible only if he is left-facing.
- Wizard `i + 1` was already visible from position `i` only if he is right-facing.

Therefore:

```text
a[i + 1] - a[i] = isLeft(i) - isRight(i + 1)
```

So the difference between adjacent values must be `-1`, `0`, or `1`.

- If `a[i + 1] = a[i] + 1`, then wizard `i` must be left-facing and wizard `i + 1` must also be left-facing.
- If `a[i + 1] = a[i] - 1`, then wizard `i` must be right-facing and wizard `i + 1` must also be right-facing.
- If `a[i + 1] = a[i]`, then the two directions must be opposite.
- Any larger absolute difference is impossible.

This leaves at most two possible full assignments: choose the direction of wizard `1`, then propagate every later direction using the adjacent rules above.

For each of the two starting choices, rebuild the whole assignment and verify it directly. The verification computes a suffix count of right-facing wizards and scans left to right with a prefix count of left-facing wizards. At position `i`, the expected value is:

```text
prefixLeft + 1 + suffixRightAfterI
```

If this equals `a[i]` for every position, the assignment is valid.

The answer is the number of valid assignments among the two starts, so it is always `0`, `1`, or `2`. The statement asks for the answer modulo `676767677`, but this modulo does not affect the implementation.

Complexity is `O(n)` per test case, with `O(n)` extra memory for the suffix array.
