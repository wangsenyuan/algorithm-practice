# E - Addition and Multiplication 2

[Problem link](https://atcoder.jp/contests/abc257/tasks/abc257_e)

**Contest:** [AtCoder Beginner Contest 257](https://atcoder.jp/contests/abc257)

time limit: 2 sec

memory limit: 1024 MiB

score: 500 points

Takahashi has an integer x, initially x = 0.

He may perform the following operation any number of times:

- Choose i with 1 <= i <= 9, pay C_i yen, and replace x with 10x + i.

Takahashi has a budget of N yen. Find the maximum possible final value of x without exceeding the
budget.

## Constraints

- 1 <= N <= 10^6
- 1 <= C_i <= N
- All input values are integers

## Input

```text
N
C_1 C_2 ... C_9
```

## Output

Print the maximum possible final value of x.

The answer may not fit in a 64-bit integer.

## Sample Input 1

```text
5
5 4 3 3 2 5 3 5 3
```

## Sample Output 1

```text
95
```

Operations i = 9 then i = 5 give 0 -> 9 -> 95, costing C_9 + C_5 = 3 + 2 = 5 yen.
No sequence within budget can produce 96 or greater.

## Sample Input 2

```text
20
1 1 1 1 1 1 1 1 1
```

## Sample Output 2

```text
99999999999999999999
```

With unit cost for every digit, the best answer is twenty 9s.

## Solution

The answer can have up to `10^6` digits, so we should build it as a string, not as an integer.

For positive integers, length is more important than digit values: any number with more digits is
larger than every number with fewer digits. Therefore the first goal is to maximize the number of
digits.

Let `best` be the largest digit among the cheapest digits. If the minimum cost is `C_best`, then the
maximum possible length is:

```text
cnt = N / C_best
```

No answer can have more than `cnt` digits, because every digit costs at least `C_best`. The code first
fills all `cnt` positions with `best`, which gives a valid maximum-length baseline.

After the length is fixed, the largest number is the lexicographically largest digit string. So we scan
positions from left to right and greedily try to make the current digit as large as possible.

At position `i`:

1. `pref` is the cost already spent on previous chosen digits.
2. `suf` is the minimum cost needed to fill the remaining suffix with `best`.
3. Try digits from `9` down to `best + 1`.
4. Choose the first digit `d` such that `pref + C_d + suf <= N`.
5. If no larger digit fits, keep `best`.

This preserves the maximum length while making each position as large as possible.

### Correctness

First, the algorithm creates a number with the maximum possible length. Every digit costs at least
`C_best`, so any number with `cnt + 1` digits would cost more than `N`. Since `cnt` copies of `best`
cost at most `N`, length `cnt` is achievable and optimal.

Now consider numbers with this fixed optimal length. Among equal-length positive integers, the larger
number is exactly the one that is lexicographically larger.

When processing a position, the algorithm tries digits from `9` downward and checks whether the rest
of the positions can still be filled at minimum cost. If a digit passes this check, then there exists a
valid completion with that digit. Any larger digit failed the same feasibility check, so no valid
maximum-length answer can put a larger digit at this position while keeping the already fixed prefix.

Thus each position is chosen as large as possible under the fixed optimal prefix. By the lexicographic
order rule, the final string is the largest possible number among all maximum-length numbers, and
therefore the largest possible final value overall.

### Complexity

Let `L = N / min(C_i)` be the answer length. For each of the `L` positions, the code checks at most
`9` digits, so the time complexity is `O(9L) = O(L)`.

The output itself has length `L`, so the memory usage is `O(L)`.
