# E. New Year and Three Musketeers

[Problem link](https://codeforces.com/problemset/problem/611/E)

time limit per test: 2 seconds

memory limit per test: 256 megabytes

input: standard input

output: standard output

Do you know the story about the three musketeers? Anyway, you must help them
now.

Richelimakieu is a cardinal in the city of Bearis. He found three brave
warriors and called them the three musketeers. Athos has strength `a`, Borthos
strength `b`, and Caramis has strength `c`.

The year 2015 is almost over and there are still `n` criminals to be defeated.
The `i`-th criminal has strength `t_i`. It is hard to defeat strong criminals:
maybe musketeers will have to fight together to achieve it.

Richelimakieu will coordinate the musketeers' actions. In each hour each
musketeer can either do nothing or be assigned to one criminal. Two or three
musketeers can be assigned to the same criminal and then their strengths are
summed up. A criminal can be defeated in exactly one hour, also if two or three
musketeers fight him. Richelimakieu cannot allow a situation where a criminal
has strength bigger than the sum of strengths of the musketeers fighting him,
because the criminal would win then.

In other words, there are three ways to defeat a criminal.

- A musketeer of strength `x` can defeat in one hour a criminal of strength not
  greater than `x`.
- Two musketeers can fight together and defeat in one hour a criminal of
  strength not greater than the sum of their strengths. The remaining third
  musketeer can either do nothing or fight some other criminal.
- All three musketeers can fight together and defeat in one hour a criminal of
  strength not greater than `a + b + c`.

Richelimakieu does not want the musketeers to fight during New Year's Eve. He
must coordinate their actions to minimize the number of hours needed to defeat
all criminals.

Find the minimum number of hours to defeat all criminals. If the musketeers
cannot defeat all criminals, print `-1`.


## Solution

Sort the three musketeer strengths:

```text
a <= b <= c
```

Also sort the criminals by strength.

If some criminal is stronger than `a + b + c`, then even all three musketeers
together cannot defeat him, so the answer is `-1`.

Otherwise, binary search the minimum number of hours. For a fixed number
`hours`, we only need to decide whether all criminals can be scheduled within
that many hours.

### Forced Hour Types

For a very strong criminal, only certain groups can defeat him. This lets us
count some mandatory hour types.

Let:

```text
ab = a + b
ac = a + c
bc = b + c
```

The strongest criminals fall into these categories:

1. `t > bc`

   Only `a+b+c` can defeat such a criminal. Each such criminal consumes one
   whole hour with all three musketeers.

2. `ac < t <= bc`

   Only pair `b+c` can defeat such a criminal without using all three. In an
   optimal schedule for a fixed `hours`, assign `b+c` to him and leave `a` free
   to defeat another smaller criminal if possible.

3. `max(ab, c) < t <= ac`

   Such a criminal cannot be defeated by `c` alone or by `a+b`, but can be
   defeated by `a+c`. This consumes pair `a+c` and leaves `b` free.

4. `c < t <= max(ab, c)`

   Such a criminal cannot be defeated by a single musketeer, but can be defeated
   by `a+b` when `a+b` is large enough. This consumes pair `a+b` and leaves `c`
   free.

The code counts these groups with `countGreater`:

```go
needAll := gt(bc)
needBC  := gt(ac) - gt(bc)
needAC  := gt(max(ab, c)) - gt(ac)
needAB  := gt(c) - gt(max(ab, c))
```

If `needAll > hours`, the schedule is impossible.

After spending `needAll` full-team hours, the remaining number of hours is:

```text
remainingHours = hours - needAll
```

The forced pair hours need:

```text
needPair = needBC + needAC + needAB
```

If `needPair > remainingHours`, the schedule is impossible.

### What Remains

After the forced categories are assigned, all remaining criminals have strength
at most `c`.

We have:

```text
freeHours = remainingHours - needPair
```

These are hours not already committed to a strong criminal.

A free hour can be used in one of two meaningful ways:

1. Three single fights:

   ```text
   a + b + c
   ```

2. One `a+b` pair fight and one `c` single fight:

   ```text
   (a+b) + c
   ```

Let:

```text
h = number of free hours used as (a+b) + c
```

Then:

```text
freeHours - h = number of free hours used as a + b + c
```

So all uncertainty is reduced to choosing one integer `h`.

### Counting Available Slots by Strength

For the remaining criminals, it is enough to compare counts by thresholds.

For any threshold `th`, every criminal with strength `> th` must be assigned to
a slot whose capacity is also `> th`.

The only relevant thresholds are:

```text
-1, a, b, c, a+b
```

because all possible slot capacities are among:

```text
a, b, c, a+b
```

For a threshold `th`, define:

```text
need(th) = number of remaining criminals with strength > th
```

We need:

```text
available_slots_with_capacity_>th >= need(th)
```

The code computes this inequality as a linear constraint on `h`.

Before the free hours, forced pair hours already create these leftover single
slots:

- each `b+c` hour leaves one `a` slot;
- each `a+c` hour leaves one `b` slot;
- each `a+b` hour leaves one `c` slot.

The free hours contribute:

- if the hour is `a+b+c`, it gives one `a`, one `b`, and one `c` slot;
- if the hour is `(a+b)+c`, it gives one `a+b` slot and one `c` slot.

Therefore, for a fixed threshold `th`, the number of usable slots can be written
as:

```text
coef * h + beta
```

The implementation builds `beta` and `coef`:

```go
if a > th {
    beta += needBC + freeHours
    coef--
}
if b > th {
    beta += needAC + freeHours
    coef--
}
if c > th {
    beta += needAB + freeHours
}
if ab > th {
    coef++
}
```

Explanation:

- `needBC` contributes leftover `a` slots from forced `b+c` hours.
- `needAC` contributes leftover `b` slots from forced `a+c` hours.
- `needAB` contributes leftover `c` slots from forced `a+b` hours.
- Every free hour initially contributes `a` and `b` as if it were a
  `a+b+c` hour.
- Changing one free hour into `(a+b)+c` removes one `a` slot and one `b` slot,
  hence `coef--` for each of them.
- The same change adds one `a+b` slot, hence `coef++` if `a+b > th`.
- The `c` slot exists in both free-hour forms, so it contributes to `beta` and
  does not depend on `h`.

For every threshold we impose:

```text
coef * h + beta >= need(th)
```

This either tightens the lower bound of `h`, tightens the upper bound of `h`, or
directly fails if `coef = 0` and `beta` is too small.

At the end, if the valid interval intersects:

```text
0 <= h <= freeHours
```

then `hours` is feasible.

### Binary Search

Feasibility is monotonic: if all criminals can be defeated in `x` hours, then
they can also be defeated in any larger number of hours.

So we binary search the smallest feasible `hours`.

The complexity is:

```text
O(N log N + log N * log N)
```

The first term is sorting criminals. Each feasibility check uses binary searches
inside the sorted array for only a constant number of thresholds.
