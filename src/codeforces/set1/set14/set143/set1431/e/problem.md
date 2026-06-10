# E. Chess Match

[Problem link](https://codeforces.com/problemset/problem/1431/E)

time limit per test: 2 seconds

memory limit per test: 512 megabytes

input: stdin

output: stdout

The final of Berland Chess Team Championship is going to be held soon. Two
teams consisting of `n` chess players each will compete for first place in the
tournament. The skill of the `i`-th player in the first team is `a_i`, and the
skill of the `i`-th player in the second team is `b_i`.

The match is held as follows: each player of the first team plays one game
against one player from the second team, so every player has exactly one
opponent. Formally, if player `i` from the first team opposes player `p_i` from
the second team, then `[p_1, p_2, ..., p_n]` is a permutation (each integer from
`1` to `n` appears exactly once).

Whenever two players of almost equal skill play a game, it will likely result in
a tie. Chess fans do not like ties, so the organizers should distribute the
players in such a way that ties are unlikely.

Let the unfairness of the match be:

```text
min over i=1..n of |a_i - b_{p_i}|
```

Assign each player from the first team an opponent from the second team so that
the unfairness is as large as possible (the greater it is, the smaller the
probability of ties is).

## Input

The first line contains an integer `t` (`1 <= t <= 3000`) — the number of test
cases.

Each test case consists of three lines:

- one line with one integer `n` (`1 <= n <= 3000`) — the number of players in
  each team;
- one line with `n` integers `a_1, a_2, ..., a_n`
  (`1 <= a_1 <= a_2 <= ... <= a_n <= 10^6`) — skills of the first team;
- one line with `n` integers `b_1, b_2, ..., b_n`
  (`1 <= b_1 <= b_2 <= ... <= b_n <= 10^6`) — skills of the second team.

It is guaranteed that the sum of `n` over all test cases does not exceed `3000`.

## Output

For each test case, print `n` integers `p_1, p_2, ..., p_n` on a separate line.
All integers from `1` to `n` must appear exactly once. The value of
`min over i=1..n of |a_i - b_{p_i}|` should be maximum possible. If there are
multiple answers, print any of them.

## Example

### Input

```text
4
4
1 2 3 4
1 2 3 4
2
1 100
100 101
2
1 100
50 51
5
1 1 1 1 1
3 3 3 3 3
```

### Output

```text
3 4 1 2
1 2
2 1
5 4 2 3 1
```

## Solution

Both arrays are sorted. The goal is to maximize:

```text
min_i |a_i - b_{p_i}|
```

### Cyclic shift

An optimal matching always comes from a cyclic shift of `b`:

```text
a[0]   ↔ b[s]
a[1]   ↔ b[s+1]
...
a[n-1] ↔ b[s+n-1]   (indices mod n)
```

Equivalently, output `p_i = (s + i) mod n + 1` for some shift `s`.

So instead of checking all `n!` permutations, try only `s = 0..n-1`. For one
shift, the unfairness is:

```text
score(s) = min_i |a[i] - b[(i+s) % n]|
```

The answer is `max_s score(s)`.

### Why an optimal matching is a cyclic shift

This subsection explains the claim above.

#### What we prove

Define:

```text
score(p) = min_i |a[i] - b[p[i]-1]|
```

The claim is:

```text
max over all permutations p of score(p)
  = max over shifts s of min_i |a[i] - b[(i+s) % n]|
```

Equivalently, for a fixed target gap `x`:

> If **any** matching has every pair with gap at least `x`, then **some cyclic
> shift** also has every pair with gap at least `x`.

That is enough: if `X*` is the optimal unfairness, a matching exists for
`x = X*`, so a shift exists for `x = X*`, and the best shift achieves `X*`.

#### Think on a circle, not a line

On a line, “no crossing” means: if `i < j`, then `p(i) < p(j)`. For a
permutation, that forces the identity `p(i) = i`.

But the optimum is often not identity. Example:

```text
a = [1, 5],  b = [3, 4]
identity:  |1-3|=2, |5-4|=1  →  score 1
shift 1:   |1-4|=3, |5-3|=2  →  score 2
```

The better matching pairs `a[0] ↔ b[1]` and `a[1] ↔ b[0]`. On a line that
crosses; on a circle of length 2 the `b`-indices `1, 0` are **consecutive mod
n**.

The correct global structure is:

```text
p(i+1) = p(i) + 1   (mod n)
```

If this holds for all `i`, then `p(i) = s + i (mod n)` for some shift `s`.

#### Local structure of a non-shift matching

Fix `x` and suppose a valid matching is **not** a shift. Then for some `i`:

```text
a[i]   ↔ b[k]
a[i+1] ↔ b[l]      with l ≠ k+1 (mod n)
```

So two consecutive `a`-players are matched to two `b`-positions that are **not
neighbors on the circle**. Since `a[i] <= a[i+1]` and `b` is sorted, there is at
least one unused `b` between them on the circle.

#### Local exchange

Compare two ways to pair `{a[i], a[i+1]}` with `{b[k], b[l]}` where `k < l`:

```text
Keep:  (a[i], b[k]),   (a[i+1], b[l])
Swap:  (a[i], b[l]),   (a[i+1], b[k])
```

Because `a[i] <= a[i+1]` and `b[k] <= b[l]`, sending the smaller `a` to the
smaller `b` does not decrease the **minimum** of the two gaps. The non-crossing
pairing on the line is locally at least as good.

Globally, starting from any valid matching for gap `x`, we can repeatedly repair
places where `p(i+1) ≠ p(i)+1 (mod n)` by rerouting along the circle without
creating any pair with gap `< x`. This process ends at:

```text
p(i) = s + i (mod n)
```

So a valid shift exists.

#### Cut-and-pair intuition

Place `b[0], ..., b[n-1]` on a circle. Pick a cut position `s`, read `b`
clockwise starting after the cut, and pair with `a[0], a[1], ...` in order.

Every shift is exactly “choose one cut, then zip with `a`”. Since `a` is sorted
forward, walking through `b` in cyclic order keeps the **worst** gap as large as
possible; zig-zagging `b`-indices while `a` moves forward tends to create one
small gap.

#### Worked example

```text
a = [1, 5],  b = [3, 4],  x = 2
```

For `a[0]=1`, any `b` in `(1-2, 1+2) = (-1, 3)` is forbidden, so only `b[1]=4`
works. Then `a[1]=5` must take `b[0]=3`:

```text
1 ↔ 4,  5 ↔ 3
```

That is shift `s=1`. No other permutation reaches `x=2`.

#### What is **not** true

- **Identity** `p(i)=i` is usually not optimal.
- **Linear non-crossing** matchings are not enough; the best matching may cross
  on a line.
- **Greedy local assignment** can fail globally (e.g. pairing `1↔3` first leaves
  score `1` instead of `2` above).

The correct structure is only: `p(i) = s + i (mod n)`.

### Binary search on the answer

Let `x` be the target unfairness. Feasibility is monotone: if some shift achieves
`score(s) >= x`, then every `y <= x` is also feasible.

`check(x)` returns whether **no** shift works (so binary search finds the first
failing `x`, and `x - 1` is optimal):

1. For each shift `s`, test whether `|a[i] - b[(i+s) % n]| >= x` for all `i`.
2. If any shift passes, `x` is feasible.
3. Binary search on `x` up to
   `max(|a[n-1] - b[0]|, |a[0] - b[n-1]|) + 1`.

### Build the permutation

After finding optimal `x`, take the first shift `s` with `score(s) >= x` and
output:

```text
p_i = (s + i) mod n + 1
```

Any such shift is accepted.

### How the code maps to this

1. Binary search the maximum feasible `x` using the shift-based `check`.
2. Scan shifts again and write the first valid permutation.
3. Print the result.

Directly scanning all shifts and taking the best `score(s)` also works in
`O(n^2)`; binary search adds a `log V` factor but keeps the same structure.

### Complexity

- One shift check: `O(n)`
- Binary search with `n` shifts per step: `O(n^2 log V)`
- Construction: `O(n^2)`
- Total per test case: `O(n^2 log V)`, fine for `sum n <= 3000`
