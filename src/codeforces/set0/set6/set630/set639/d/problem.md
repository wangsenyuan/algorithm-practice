# Problem

Codeforces is a wonderful platform, and one of its features shows how much someone contributes to the community. Every registered user has a contribution — an integer value, not necessarily positive. There are $n$ registered users, and the $i$-th of them has contribution $t_i$.

Limak is a little polar bear, and he is new to competitive programming. He does not even have an account on Codeforces, but he is able to upvote existing blogs and comments. Assume that every registered user has infinitely many blogs and comments.

Limak can spend $b$ minutes to read one blog and upvote it. The author's contribution increases by $5$.
Limak can spend $c$ minutes to read one comment and upvote it. The author's contribution increases by $1$.

Note that it is possible that Limak reads blogs faster than comments.

Limak likes ties. He thinks it would be awesome to see a tie between at least $k$ registered users. To make it happen, he is going to spend some time reading and upvoting. After that, there should exist an integer value $x$ such that at least $k$ registered users have contribution exactly $x$.

How much time does Limak need to achieve his goal?

## Input

The first line contains four integers $n$, $k$, $b$, and $c$ ($2 \le k \le n \le 200000$, $1 \le b, c \le 1000$) — the number of registered users, the required minimum number of users with the same contribution, time needed to read and upvote a blog, and time needed to read and upvote a comment, respectively.

The second line contains $n$ integers $t_1, t_2, \ldots, t_n$ ($|t_i| \le 10^9$), where $t_i$ denotes the contribution of the $i$-th registered user.

## Output

Print the minimum number of minutes Limak will spend to get a tie between at least $k$ registered users.

## Examples

### Example 1

**Input**

```
4 3 100 30
12 2 6 1
```

**Output**

```
220
```

### Example 2

**Input**

```
4 3 30 100
12 2 6 1
```

**Output**

```
190
```

### Example 3

**Input**

```
6 2 987 789
-8 42 -4 -65 -8 -8
```

**Output**

```
0
```

## Note

In the first sample, there are $4$ registered users and Limak wants a tie between at least $3$ of them. Limak can behave as follows:

- Spend $100$ minutes to read one blog of the $4$-th user and increase contribution from $1$ to $6$.
- Spend $4 \cdot 30 = 120$ minutes to read four comments of the $2$-nd user and increase contribution from $2$ to $6$ (four increments by $1$).

In this scenario, Limak spends $100 + 4 \cdot 30 = 220$ minutes, and then users $2$, $3$, and $4$ each have contribution $6$.

In the second sample, Limak needs $30$ minutes per blog and $100$ minutes per comment. This time he can get $3$ users with contribution equal to $12$ by spending $100 + 3 \cdot 30 = 190$ minutes:

- Spend $2 \cdot 30 = 60$ minutes to read two blogs of the $2$-nd user, increasing contribution from $2$ to $12$.
- Spend $30 + 100$ minutes to read one blog and one comment of the $3$-rd user. Their contribution changes from $6$ to $6 + 5 + 1 = 12$.


### ideas
1. 假设目标分数是x，需要在最少的时间，去得到至少k个人的分数是x
2. 因为comment的增量是1，所以在给给定x的情况下 x >= t, 肯定是可以使的t增加到x
3. 如果b <= 5 * c, 那么优先使用操作1，(x - t) % 5 * c + (x - t) / 5 * b (花费时间最少)
4. 如果 b > 5 * c, 那么就是 （x - t) * c (全部用操作2)
5. 如果 b >= 5 * c, 那么就简单的一个前缀dp，就可以了（肯定是某个t[i], 然后前缀k个右边的得到分数t[i])的最优解
6. 现在考虑b < 5 * c, 这个有点麻烦

## Detailed Explanation

We want at least `k` users to finish with the same contribution `x`.

For a fixed target `x`, every user with `t_i > x` is useless, because we can only increase contributions.
Every user with `t_i <= x` can be raised to `x`, and we want the cheapest `k` of them.

So the whole problem is:

1. choose a target value `x`,
2. compute the minimum cost to raise at least `k` users to exactly `x`,
3. minimize over all `x`.

The difficult part is understanding which targets must be checked and how to compute the cheapest `k` users fast.

### 1. Cost of raising one user to `x`

Let:

- blog: `+5` cost `b`
- comment: `+1` cost `c`

If we need to add `d = x - t_i`, then:

- one option is to use only comments: cost `d * c`
- another is to use some blogs plus a few comments

If `b >= 5*c`, then one blog is never better than five comments.
So the optimal cost is simply:

- `cost(d) = d * c`

That case is easy.

If `b < 5*c`, then blogs are useful.
For fixed `d`, the best strategy is:

- use as many blogs as possible,
- then use comments for the remainder modulo `5`

So:

- `cost(d) = floor(d / 5) * b + (d mod 5) * c`

This can also be written more structurally once we look at `x mod 5`.

### 2. The easy case: `b >= 5*c`

If comments are always at least as good as blogs, then for fixed `x`:

- cost to raise `t_i` to `x` is `(x - t_i) * c`

If we sort contributions:

- `t_1 <= t_2 <= ... <= t_n`

then for a target `x = t_r`, the cheapest `k` users are exactly the `k` largest values among those `<= x`, i.e. the window:

- `t_{r-k+1}, ..., t_r`

because larger starting values are cheaper to raise.

So we just slide a window of length `k` and compute:

- `(k * t_r - sum(window)) * c`

This is what `solveCommentsOnly` does.

### 3. Why the same idea fails when `b < 5*c`

This is the subtle part.

When blogs are useful, the cost is not determined only by how far `t_i` is below `x`.
It also depends on the remainder modulo `5`.

Example:

- `b = 3`, `c = 10`
- target `x = 30`

Compare:

- `t = 29`, need `1`, cost `10`
- `t = 25`, need `5`, cost `3`

Although `25` is farther away than `29`, it is cheaper to raise, because it fits one blog exactly.

So for `b < 5*c`, the cheapest `k` users for a fixed `x` are **not** necessarily the `k` largest contributions `<= x`.

That is why the previous simpler approach breaks on extra tests.

### 4. Group by target residue modulo `5`

Now fix:

- `x = 5a + y`, where `y in {0,1,2,3,4}`

Take one user with contribution `t`.
Write:

- `t = 5q + r`, where `r in {0,1,2,3,4}`

If we want to raise `t` to `x`, then:

- if `r <= y`, we need exactly `a - q` blogs and `y - r` comments
- if `r > y`, then we need `a - q - 1` full fives from the difference, but because of the borrow effect it becomes
  one extra blog and `(y - r + 5)` comments

It is cleaner to encode this as:

- the user becomes available only when `a >= need`
- and once available, their cost is linear in `a`

More precisely:

- `need = q`, if `r <= y`
- `need = q + 1`, if `r > y`

and for all `a >= need`,

- `cost = a * b - weight`

where `weight` is a constant depending only on `(q, r, y)`:

- `weight = q*b - (y-r)c`, if `r <= y`
- `weight = q*b - (y-r+5)c + b`, if `r > y`

In the code this is written as:

- `need`
- `weight`

and the total cost for chosen users becomes:

- `k * a * b - sum(weight)`

So for fixed residue `y` and fixed threshold `a`, we want:

- among all users with `need <= a`,
- choose `k` users with maximum `weight`

because that minimizes total cost.

### 5. Sweep over `a`

For one fixed residue `y`, each user contributes one item:

- available starting from threshold `need`
- value `weight`

Sort all items by `need`.

Now sweep `a` from small to large, but only at the critical values that actually appear as some `need`.
As we sweep:

- add every newly available user's `weight`
- maintain the best `k` weights among all available users

This is done with a min-heap:

- keep exactly the `k` largest weights seen so far
- if heap size exceeds `k`, pop the smallest

Let `sum` be the sum of the `k` kept weights.
Then current answer candidate is:

- `k * a * b - sum`

We do this for all `y = 0..4`, and take the minimum.

### 6. Why only `5` residue classes are needed

Because blogs change contribution by `5`.
So once `x mod 5` is fixed, every user's optimal decomposition into blogs/comments is determined only by:

- their own `t mod 5`
- the quotient part `floor(t / 5)`

That collapses the infinitely many possible targets into:

- five independent sweeps over thresholds `a`

### 7. Complexity

For each residue `y`:

1. build `n` items,
2. sort them by `need`,
3. sweep once with a heap.

So the total complexity is:

- `O(5 * n log n)`

which is easily fast enough for `n <= 2 * 10^5`.

### 8. Summary

The solution has two cases:

1. `b >= 5*c`
   - blogs are useless,
   - answer is a simple sliding-window minimum on the sorted array.

2. `b < 5*c`
   - blogs matter,
   - cost depends on residues modulo `5`,
   - for each target residue `y`, transform every user into:
     - a threshold `need`
     - a value `weight`
   - sweep thresholds and keep the best `k` available weights with a heap.

The most important insight is:

- for `b < 5*c`, “closest contributions to `x`” is **not** the same as “cheapest to raise to `x`”,
- because the remainder modulo `5` changes how many expensive `+1` operations are needed.
