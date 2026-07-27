# B. Decidophobia

[Problem link](https://codeforces.com/problemset/problem/2239/B)

**Contest:** [Codeforces Round 1105 (Div. 1)](https://codeforces.com/contest/2239)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

There are `n` people attending a round-table party, numbered `1, 2, 3, ..., n` in
clockwise order. You have prepared some gifts to distribute among them.

Each person `i` has a weight `a_i` and a common field of view `d`. The field of view for
person `i` consists of the `d` people sitting clockwise and the `d` people sitting
counter-clockwise from them (a total of `2d` people excluding person `i`).

The happiness gained by person `i` is determined by the following rules:

- If person `i` receives a gift, and there are `x` people within their field of view who
  did not receive a gift, they gain `x * a_i` happiness.
- If person `i` does not receive a gift, and there are `x` people within their field of
  view who received a gift, they incur `-x * a_i` happiness.

You want to maximize the total happiness of all `n` people combined. Find this maximum
value.

## Input

Each test contains multiple test cases. The first line contains the number of test cases
`t` (`1 <= t <= 10^4`). The description of the test cases follows.

The first line of each test case contains two integers `n` and `d`
(`3 <= n <= 2 * 10^5`, `1 <= d < n / 2`).

The second line contains `n` integers `a_1, a_2, ..., a_n` (`1 <= a_i <= 10^8`), where
`a_i` is the weight of the `i`-th person.

It is guaranteed that the sum of `n` over all test cases does not exceed `10^6`.

## Output

For each test case, output a single integer representing the maximum total happiness.

## Example

### Input

```text
5
3 1
1 2 3
5 1
1 4 5 2 6
6 2
1 1 4 5 1 4
10 2
230 24 3 42 432 234 934 2389 333 444
3 1
100000000 100000000 100000000
```

### Output

```text
3
15
26
8590
0
```

## Note

In the first test case, there are 3 people sitting in a circle. For each person `i`, the
field of view is `d = 1`, which includes the 1 person sitting clockwise and the 1 person
sitting counter-clockwise from them. If person 2 receives the gift, they gain happiness
from 2 neighbors who did not receive it, resulting in `2 * a_2 = 2 * 2 = 4`. However,
person 1 and person 3 incur loss because they did not receive a gift but have a neighbor
who did. After calculating all possibilities, the maximum total happiness is 3.

In the second test case, `n = 5`, `d = 1`. The best solution is to give gifts to persons
2, 3, and 5, which can achieve the maximum total happiness value of 15.


## ideas
1. 假设a[i]被选中了, 它前后都没有选中的情况下, 它的收益 = 2 * d * a[i] - (sum(a[l...r]) - a[i])
2. = a[i] + a[i] - a[l] + a[i] - a[l+1] + ... a[i] - a[r]
3. 这是在[l...r]中间只选一个的情况
4. 考虑选择两个情况 a[i1], a[i2], a[i1] > a[i2]
5. 好像有点麻烦的, 但是感觉选择两个应该不会更优
6. a[i2]的选择, 造成a[i1]少了一点收益

## Solution

### Setup

People `0..n-1` sit in a circle. Let

```text
s_i ∈ {0, 1}   = 1 iff person i gets a gift
FOV(i)         = the 2d neighbors within distance d (not including i)
```

Person `i`'s happiness:

- if `s_i = 1`: `+ a_i · (# of FOV neighbors with s_j = 0)`
- if `s_i = 0`: `- a_i · (# of FOV neighbors with s_j = 1)`

### Step 1 — Simplify one person's happiness

Let `S_i = Σ_{j ∈ FOV(i)} s_j` (number of gifted neighbors).

```text
h_i = s_i · a_i · (2d - S_i)  +  (1 - s_i) · (-a_i · S_i)
    = a_i · [ 2d·s_i - s_i·S_i - S_i + s_i·S_i ]
    = a_i · (2d·s_i - S_i)
```

The `± s_i·S_i` terms cancel. So:

```text
h_i = a_i · (2d · s_i - Σ_{j ∈ FOV(i)} s_j)
```

### Step 2 — Total happiness

```text
H = Σ_i h_i
  = Σ_i a_i · (2d · s_i - Σ_{j ∈ FOV(i)} s_j)
  = 2d · Σ_i a_i s_i  -  Σ_i a_i · Σ_{j ∈ FOV(i)} s_j
```

Swap the double sum in the second term: each `s_j` is multiplied by the sum of `a_i`
over all `i` that can "see" `j`.

On a circle with symmetric FOV, `j ∈ FOV(i)` iff `i ∈ FOV(j)`, so

```text
Σ_i a_i · Σ_{j ∈ FOV(i)} s_j  =  Σ_j s_j · A_j
```

where

```text
A_j = Σ_{i ∈ FOV(j)} a_i
```

is just the sum of weights of `j`'s neighbors.

Therefore:

```text
H = Σ_j s_j · (2d · a_j - A_j)
```

### Step 3 — Why choices are independent

`H` is a **sum of terms each depending on a single `s_j`**. There is no `s_j · s_k`
product left.

So for each `j` separately:

```text
if 2d·a_j - A_j > 0  → set s_j = 1 (add that value)
if 2d·a_j - A_j < 0  → set s_j = 0
if equal to 0        → either is fine
```

```text
answer = Σ_j max(0, 2d·a_j - A_j)
```

Intuition: giving `j` a gift contributes `+2d·a_j` from `j`'s own formula, but also
subtracts `a_i` once for every neighbor `i` that sees `j` (because each such neighbor's
formula picks up `s_j`). Net coefficient of `s_j` is exactly `2d·a_j - A_j`.

### Step 4 — Computing `A_j` in O(n)

`A_j` = sum of `a` on `[j-d, j)` plus `(j, j+d]` on the circle.

Double the array and build a prefix sum:

```text
pref[0] = 0
pref[k+1] = pref[k] + a[k mod n]   for k = 0 .. 2n-1
```

Then for each `i`:

```text
A_i = sum(i-d .. i-1) + sum(i+1 .. i+d)
    = rangeSum(i+n-d, i+n) + rangeSum(i+1, i+1+d)
contrib = 2·d·a_i - A_i
answer += max(0, contrib)
```

Total time `O(n)` per test, fine under `Σ n ≤ 10⁶`.

### Sample 1 check

`n=3`, `d=1`, `a=[1,2,3]`. Each FOV is the other two people.

| i | A_i | 2d·a_i - A_i | take? |
|---|-----|--------------|-------|
| 0 | 2+3=5 | 2-5=**-3** | no |
| 1 | 1+3=4 | 4-4=**0** | no |
| 2 | 1+2=3 | 6-3=**3** | yes |

Answer `3` (gift only person 3). Matches the brute-force optimum.

### Why the "interaction worry" is unnecessary

Giving two people gifts looks coupled (each reduces the other's ungifted-neighbor
count). In the algebra those couplings appear twice with opposite signs and cancel, so
the net score really is additive over chosen people.