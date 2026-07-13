# D - Iroha and Haiku (New ABC Edition)

[Problem link](https://atcoder.jp/contests/abc265/tasks/abc265_d)

**Contest:** [AtCoder Beginner Contest 265](https://atcoder.jp/contests/abc265)

time limit: 2 sec

memory limit: 1024 MiB

score: 400 points

There is a sequence `A = (A_0, ..., A_{N-1})` of length `N`.

Determine whether there exists a tuple of integers `(x, y, z, w)` that satisfies
all of the following:

- `0 <= x < y < z < w <= N`
- `A_x + A_{x+1} + ... + A_{y-1} = P`
- `A_y + A_{y+1} + ... + A_{z-1} = Q`
- `A_z + A_{z+1} + ... + A_{w-1} = R`

## Constraints

- `3 <= N <= 2 * 10^5`
- `1 <= A_i <= 10^9`
- `1 <= P, Q, R <= 10^15`
- All values in the input are integers

## Input

```text
N P Q R
A_0 A_1 ... A_{N-1}
```

## Output

Print `Yes` if such a tuple exists; otherwise print `No`.

## Sample Input 1

```text
10 5 7 5
1 3 2 2 2 3 1 4 3 2
```

## Sample Output 1

```text
Yes
```

`(x, y, z, w) = (1, 3, 6, 8)` satisfies the conditions.

## Sample Input 2

```text
9 100 101 100
31 41 59 26 53 58 97 93 23
```

## Sample Output 2

```text
No
```

## Sample Input 3

```text
7 1 1 1
1 1 1 1 1 1 1
```

## Sample Output 3

```text
Yes
```

## Solution

Use prefix sums and process the three target segments **from right to left**
(`R`, then `Q`, then `P`).

### Key idea

Let `S[i] = A_0 + ... + A_{i-1}` (`S[0] = 0`). Then a subarray sum
`A[l] + ... + A[r-1] = S[r] - S[l]`.

Because `A_i >= 1`, prefix sums are strictly increasing, so each sum appears at
most once.

Work backwards:

1. Find every index `z` such that some `w > z` has `S[w] - S[z] = R`.
2. Find every index `y` such that some `z > y` has `S[z] - S[y] = Q` **and**
   `z` is valid from step 1.
3. Find every index `x` such that some `y > x` has `S[y] - S[x] = P` **and**
   `y` is valid from step 2.

If any such `x` exists, answer `Yes`.

### Bit DP on endpoints

Maintain `dp[i]` as a bitset of progress ending / starting at index `i`:

| bit | meaning |
| --- | --- |
| `0` | free start (initially set on every index) |
| `1` | an `R`-segment starts at this index |
| `2` | a `Q`-segment starts here, followed by a valid `R` |
| `3` | a `P`-segment starts here, followed by valid `Q` then `R` |

One scan `play(d, w)`:

- Walk left to right, keep `pref[S[i]] = i`.
- When visiting end `i`, if `dp[i]` has bit `d`, look up `pref[S[i] - w] = j`.
- Then `A[j..i)` sums to `w`, so set bit `d+1` at `j` (the segment start).

Call:

```text
play(0, R)  // bit0 -> bit1
play(1, Q)  // bit1 -> bit2
play(2, P)  // bit2 -> bit3
```

Finally, if any `dp[i]` has bit `3`, print `Yes`; otherwise `No`.

### Complexity

Three linear passes with a hash map of prefix sums: `O(N)` expected time and
`O(N)` memory, which fits `N <= 2 * 10^5`.
