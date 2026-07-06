# Ex - Dice Sum 2

[Problem link](https://atcoder.jp/contests/abc257/tasks/abc257_h)

**Contest:** [AtCoder Beginner Contest 257](https://atcoder.jp/contests/abc257)

time limit: 2 sec

memory limit: 1024 MiB

score: 600 points

The shop "Saikoroya" sells N six-sided dice. The i-th die shows A_{i,1}, ..., A_{i,6} on its faces
and costs C_i yen.

Takahashi buys exactly K dice. During a promotion, he rolls each purchased die once and receives
money equal to the square of the sum of the rolled values. Each face appears uniformly at random,
independently across dice.

Choose K dice to maximize the expected value of

(reward) - (total purchase cost)

Print the maximized expected value modulo 998244353.

The answer is defined as the unique integer z with 0 <= z < 998244353 such that x * z ≡ y (mod
998244353), when the expected value equals y/x in lowest terms and x is not divisible by 998244353.

## Constraints

- 1 <= N <= 1000
- 1 <= K <= N
- 1 <= C_i <= 10^5
- 1 <= A_{i,j} <= 10^5
- All input values are integers

## Input

```text
N K
C_1 C_2 ... C_N
A_{1,1} A_{1,2} ... A_{1,6}
...
A_{N,1} A_{N,2} ... A_{N,6}
```

## Output

Print the answer modulo 998244353.

## Sample Input 1

```text
3 2
1 2 3
1 1 1 1 1 1
2 2 2 2 2 2
3 3 3 3 3 3
```

## Sample Output 1

```text
20
```

Buying dice 2 and 3 gives expected value (2 + 3)^2 - (2 + 3) = 20, which is optimal.

## Sample Input 2

```text
10 5
2 5 6 5 2 1 7 9 7 2
5 5 2 4 7 6
2 2 8 7 7 9
8 1 9 6 10 8
8 6 10 3 3 9
1 10 5 8 1 10
7 8 4 8 6 5
1 10 2 5 1 7
7 4 1 4 5 4
5 10 1 5 1 2
5 1 2 3 6 2
```

## Sample Output 2

```text
1014
```

## Solution

The direct object is awkward because the reward is the square of a random sum.
The first step is to rewrite the expected value so that each die contributes one
simple two-dimensional vector.

For die `i`, define:

```text
S_i = A_{i,1} + A_{i,2} + ... + A_{i,6}
Q_i = A_{i,1}^2 + A_{i,2}^2 + ... + A_{i,6}^2
```

If the random value rolled by this die is `X_i`, then:

```text
E[X_i]   = S_i / 6
E[X_i^2] = Q_i / 6
```

Suppose we choose a set `T` of exactly `K` dice. The reward before subtracting
costs is:

```text
E[(sum_{i in T} X_i)^2]
```

Expand the square:

```text
E[(sum X_i)^2]
= E[sum X_i^2 + 2 * sum_{i < j} X_i X_j]
= sum E[X_i^2] + 2 * sum_{i < j} E[X_i X_j]
```

The dice are independent, so:

```text
E[X_i X_j] = E[X_i] * E[X_j] = (S_i / 6) * (S_j / 6)
```

Therefore:

```text
E[(sum X_i)^2]
= sum (Q_i / 6) + 2 * sum_{i < j} (S_i S_j / 36)
```

Multiplying by `36` removes all denominators:

```text
36 * E[(sum X_i)^2]
= 6 * sum Q_i + 2 * sum_{i < j} S_i S_j
```

The cross term can be written using the square of the total `S`:

```text
(sum S_i)^2 = sum S_i^2 + 2 * sum_{i < j} S_i S_j
```

So:

```text
2 * sum_{i < j} S_i S_j = (sum S_i)^2 - sum S_i^2
```

Substitute this:

```text
36 * E[(sum X_i)^2]
= (sum S_i)^2 + sum (6Q_i - S_i^2)
```

The cost is not random. Subtracting total cost and still multiplying by `36`:

```text
36 * objective
= (sum S_i)^2 + sum (6Q_i - S_i^2 - 36C_i)
```

Now define one vector for each die:

```text
x_i = S_i
y_i = 6Q_i - S_i^2 - 36C_i
```

Then the problem becomes:

```text
Choose exactly K vectors to maximize:

(sum x_i)^2 + sum y_i
```

At the end, the implementation divides by `36` modulo `998244353`, using
`36^(mod-2)`.

### Why this is still hard

The term `sum y_i` is linear, but `(sum x_i)^2` depends on the total `x` of the
chosen set. This means we cannot simply sort dice independently by one fixed
score.

For example, a die with large `x` may become more valuable when the current total
`x` is already large, because the square term rewards increasing the total.

The official editorial's trick is to look at a family of easier problems.

For any real number `t`, consider the linear score:

```text
t * sum x_i + sum y_i
```

For a fixed `t`, choosing exactly `K` dice to maximize this is easy:

```text
pick the K dice with largest individual value t*x_i + y_i
```

So if we can argue that the optimal answer for `(sum x)^2 + sum y` appears among
the top-`K` sets for some value of `t`, then we only need to enumerate how that
top-`K` set changes as `t` moves.

### Why checking linear scores is enough

Think of every possible chosen set as one point:

```text
X = sum x_i
Y = sum y_i
```

The value we want to maximize is:

```text
X^2 + Y
```

Among all possible `(X, Y)` points, ignore any point that is never best for any
linear expression:

```text
tX + Y
```

Such a point is hidden below another point or below a segment between other
points. It cannot be the maximum of `X^2 + Y` either.

More concretely, if a point `(X, Y)` lies below a segment between two other
available points `(X1, Y1)` and `(X2, Y2)`, then for some `0 <= p <= 1`:

```text
X = pX1 + (1-p)X2
Y <= pY1 + (1-p)Y2
```

Because `X^2` is a convex function:

```text
X^2 <= pX1^2 + (1-p)X2^2
```

Therefore:

```text
X^2 + Y <= p(X1^2 + Y1) + (1-p)(X2^2 + Y2)
```

So at least one endpoint has value no smaller than `(X, Y)`. This means the
maximum happens on the upper convex hull of the possible `(X, Y)` points.

Every vertex of that upper hull is the best point for some linear expression
`tX + Y`. For that fixed `t`, maximizing `tX + Y` over chosen sets is the same as
choosing the `K` largest dice by `t*x_i + y_i`.

That is the reason the implementation only needs to inspect top-`K` sets that
appear while sorting dice by:

```text
t*x_i + y_i
```

### How the order changes

For two dice `i` and `j`, their relative order by `t*x+y` changes only when:

```text
t*x_i + y_i = t*x_j + y_j
```

Rearrange:

```text
t * (x_i - x_j) = y_j - y_i
t = (y_j - y_i) / (x_i - x_j)
```

So every pair of dice gives at most one crossing time. There are `O(N^2)` such
events.

The implementation stores each event as the reduced fraction:

```text
num / den = (y_j - y_i) / (x_i - x_j)
```

with `den > 0`, and sorts events by:

```text
num1 / den1 < num2 / den2
```

using cross multiplication:

```text
num1 * den2 < num2 * den1
```

This avoids floating-point precision problems.

Pairs with the same `x` never cross, because:

```text
t*x_i + y_i
```

and

```text
t*x_j + y_j
```

are parallel lines. Their relative order is fixed by `y`.

### Initial order

The code starts with dice sorted by increasing `x`.

For very negative `t`, the term `t*x` dominates. A smaller `x` gives a larger
score because `t` is negative. Therefore increasing `x` is the correct order at
the far left of the sweep.

Ties are broken by larger `y`, then by index. The tie-break only makes the order
deterministic; the correctness comes from processing all equal crossing times
together.

The current chosen set is always:

```text
the first K dice in order
```

The code maintains:

```text
sumX = sum x_i of the first K dice
sumY = sum y_i of the first K dice
```

and evaluates:

```text
sumX * sumX + sumY
```

after every event group.

### Why reversing a segment is correct

At a crossing time, several dice may have exactly the same value of `t*x+y`.

First consider the simple case: only two dice cross. Before the crossing, one is
before the other. After the crossing, their relative order is swapped. In the
sorted order these two dice are adjacent at the moment of crossing, so reversing
the segment between their current positions swaps them.

For multiple simultaneous crossings, all dice connected by equal-time crossings
belong to one group. Inside such a group, the order before and after the crossing
is reversed.

Why reversed? Suppose all dice in this group have the same score at time `t0`.
For a small positive `eps`, compare times:

```text
t0 - eps
t0 + eps
```

At `t0`, the scores are equal. The difference just before or after depends on:

```text
(t - t0) * x_i
```

So the order by `x` flips when passing through `t0`. Therefore the whole block of
equal-score dice reverses.

The implementation handles this as follows:

1. Gather all pair events with the same crossing time.
2. Use DSU to connect dice that cross at that same time.
3. For each connected component, find its current position interval `[lo, hi]`.
4. Reverse that interval in the global order.

The single-pair case is common, so the code has a fast path that reverses the
two-position interval directly.

### Updating the selected sum during a reversal

Only the boundary between selected and unselected dice matters.

The selected dice are positions:

```text
0, 1, ..., K-1
```

When reversing an interval `[l, r]`, there are three cases:

1. `r < K`: the whole interval is selected.
2. `K <= l`: the whole interval is unselected.
3. `l < K <= r`: the interval crosses the boundary.

In cases 1 and 2, the set of selected dice does not change. Their order changes,
but `sumX` and `sumY` stay the same.

Only case 3 changes the selected set.

Before reversal, the selected part inside the interval is:

```text
[l, K-1]
```

Its length is:

```text
cnt = K - l
```

After reversing `[l, r]`, those `cnt` selected slots are filled by the last `cnt`
dice of the old interval:

```text
[r-cnt+1, r]
```

So the code subtracts the old selected dice:

```text
for i := l; i < K; i++ {
    remove order[i]
}
```

and adds the new selected dice:

```text
for i := r-cnt+1; i <= r; i++ {
    add order[i]
}
```

Then it physically reverses the interval and updates `pos[id]` for every swapped
die.

### Mapping to the code

The main arrays are:

```text
x[i] = S_i
y[i] = 6Q_i - S_i^2 - 36C_i
```

`order` stores dice indices in the current decreasing order of `t*x+y`.
Because the sweep starts at very negative `t`, this starts as increasing `x`.

`pos[id]` is the inverse array:

```text
pos[dice id] = current position in order
```

Each `Event` stores:

```text
u, v     the two dice
num, den the reduced crossing time num/den
```

The event loop processes all events with the same time together:

```text
for l := 0; l < len(events); {
    r := l + 1
    for r < len(events) && sameEventTime(events[l], events[r]) {
        r++
    }
    ...
    l = r
}
```

For one event, it reverses the current interval between the two dice.

For several events at the same time, it compresses involved dice, unions the
pairs with DSU, and reverses one interval for each connected component.

After each crossing time is processed, the current first `K` dice represent one
top-`K` set for values of `t` just after that crossing, so the answer candidate is
updated:

```text
best = max(best, sumX*sumX + sumY)
```

### Correctness

First, the transformation from dice to vectors preserves the objective up to a
constant factor of `36`. For any chosen set, the code computes exactly:

```text
36 * (expected reward - total cost)
= (sum x)^2 + sum y
```

Since `36` is positive, maximizing the scaled value is equivalent to maximizing
the original expected value.

Second, for any fixed real `t`, the best set for the linear expression
`t*sumX + sumY` is obtained by choosing the `K` dice with largest individual
scores `t*x_i + y_i`. This is because the total score is just the sum of the
chosen individual scores.

Third, the maximum of `(sumX)^2 + sumY` is attained by a set that is optimal for
some linear expression `t*sumX + sumY`. Points that are never optimal for any
linear expression lie below the upper convex hull, and a convexity argument shows
that such a point cannot uniquely improve `X^2 + Y` over the hull endpoints.

Fourth, as `t` increases from negative infinity to positive infinity, the sorted
order by `t*x_i+y_i` changes only at pair crossing times. Between two consecutive
crossing times, the order is fixed, so the top-`K` set is fixed. Therefore
checking the initial set and the set after every crossing time checks all
possible linear-optimal top-`K` sets.

Finally, at a crossing time, each connected group of equal-time crossings forms a
contiguous block in the current order, and its order reverses after the crossing.
The implementation performs exactly that reversal and updates the selected sums
only when the block crosses the `K` boundary. Therefore `sumX` and `sumY` always
describe the first `K` dice of the current order, and every evaluated candidate
is correct.

Combining these facts, the maximum `best` found by the sweep is exactly the
maximum scaled objective. Multiplying by `36^{-1}` modulo `998244353` gives the
required answer.

### Complexity

There are at most:

```text
N * (N - 1) / 2
```

crossing events.

Generating them takes `O(N^2)`. Sorting them takes `O(N^2 log N)`.

During the sweep, each event belongs to exactly one event group. Reversing
segments and updating positions costs proportional to the total number of moved
positions. For this problem size, `N <= 1000`, this direct implementation is fast
enough. The dominant cost is the event sort.

The memory usage is `O(N^2)` for the event list.
