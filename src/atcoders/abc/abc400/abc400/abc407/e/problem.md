# E - Most Valuable Parentheses

[Problem link](https://atcoder.jp/contests/abc407/tasks/abc407_e)

**Contest:** [AtCoder Beginner Contest 407](https://atcoder.jp/contests/abc407)

time limit: 2 sec

memory limit: 1024 MiB

score: 450 points

You are given a sequence of non-negative integers `A = (A_1, ..., A_{2N})` of length `2N`.

Define the score of a parenthesis sequence `s` of length `2N` as follows:

- For every position `i` where the `i`-th character of `s` is `)`, set `A_i` to `0`, then take the sum of all elements of `A`.

Find the maximum possible score of a **correct** parenthesis sequence of length `2N`.

You are given `T` test cases; solve each.

A correct parenthesis sequence is a string that can be reduced to the empty string by repeatedly
removing substrings equal to `()`.

## Constraints

- `1 <= T <= 500`
- `1 <= N <= 2 * 10^5`
- The sum of `N` over all test cases in one input file is at most `2 * 10^5`.
- `0 <= A_i <= 10^9` (`1 <= i <= 2N`)
- All input values are integers.

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
A_1
A_2
...
A_{2N}
```

## Output

Print `T` lines. The `i`-th line should contain the answer for the `i`-th test case.

## Sample Input 1

```text
2
3
400
500
200
100
300
600
6
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
1000000000
```

## Sample Output 1

```text
1200
6000000000
```

### Note

In the first test case, choosing `(())()` gives score `400 + 500 + 0 + 0 + 300 + 0 = 1200`.

No correct parenthesis sequence yields a higher score.

The answer may exceed 32-bit integer range.

## Solution

The score keeps the values at positions chosen as `(` and removes the values at positions chosen as `)`.
So we want to maximize the sum of values on the `N` positions chosen as opening parentheses.

For a correct parenthesis sequence, every prefix must contain at least as many `(` as `)`. Therefore, in
the first `2t+1` positions, we must already have chosen at least `t+1` opening parentheses.

Using 0-based indices, this means:

```text
after index 0, choose at least 1 opening position
after index 2, choose at least 2 opening positions
after index 4, choose at least 3 opening positions
...
```

So every even index is a deadline by which we must choose one more `(`.

### Greedy View

Process positions from left to right.

- Odd indices do not immediately force a new opening parenthesis, so store their values as candidates.
- Even indices force us to add one more opening parenthesis right now.

At an even index `i`, there are two choices:

1. choose this even position `i` as `(` and gain `A_i`;
2. choose one previously seen odd position as `(` instead.

If we use a previous odd position, we should choose the largest available value among them. The
implementation keeps those odd-position candidates in a max-heap.

For each even index:

```text
if best previous odd value > A_i:
    choose that odd value
    put A_i into the heap as a future candidate
else:
    choose A_i
```

Why can `A_i` be put into the heap after choosing a previous odd value? Because exactly one new opening
position has been selected for this deadline. The current position was not selected now, but it has still
been seen, so it may be selected later if it is useful.

This is equivalent to repeatedly keeping the best possible set of chosen opening values after each
deadline.

### Why This Is Correct

Consider the moment we reach an even index `i`. By the parenthesis-prefix condition, the number of opening
positions chosen must increase by one.

All positions seen so far are the only positions available for satisfying this deadline. Among the
positions that have not already been fixed as chosen, picking the largest value is always optimal:

- it satisfies the same deadline as any smaller value;
- it gives at least as much score;
- choosing a smaller value cannot make future deadlines easier, because both choices consume exactly one
  available position.

Thus, at every deadline, choosing the maximum available candidate is safe. The heap stores exactly the
available candidates that can still be swapped into the chosen set. By applying this choice at all `N`
deadlines, the algorithm maximizes the total value assigned to `(` positions.

### Algorithm

1. Initialize `ans = 0` and an empty max-heap.
2. Scan `A` from left to right.
3. If the index is odd, push `A_i` into the heap.
4. If the index is even:
   - if the heap has a value larger than `A_i`, pop it, add it to `ans`, and push `A_i`;
   - otherwise add `A_i` to `ans`.
5. Print `ans`.

### Complexity

Each value is pushed or popped at most once. Heap operations cost `O(log N)`, so each test case runs in
`O(N log N)` time and uses `O(N)` memory.
