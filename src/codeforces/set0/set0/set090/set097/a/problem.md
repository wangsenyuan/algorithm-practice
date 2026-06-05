# A. Domino

[Problem link](https://codeforces.com/problemset/problem/97/A)

time limit per test: 0.5 second

memory limit per test: 256 megabytes

input: stdin

output: stdout

Little Gennady was presented with a set of domino for his birthday. The set
consists of 28 different dominoes of size `2 x 1`. Both halves of each domino
contain one digit from `0` to `6`.

```text
0-0 0-1 0-2 0-3 0-4 0-5 0-6
1-1 1-2 1-3 1-4 1-5 1-6
2-2 2-3 2-4 2-5 2-6
3-3 3-4 3-5 3-6
4-4 4-5 4-6
5-5 5-6
6-6
```

A figure consisting of 28 dominoes is called magic if it can be fully covered
with 14 non-intersecting squares of size `2 x 2` so that each square contains
four equal numbers.

Every time Gennady assembles a magic figure, some magic properties of the set
appear: he wins the next contest. Gennady noticed that he cannot assemble a
figure that has already been assembled, otherwise someone else wins the contest.

Gennady chose a checked field of size `n x m` and put there rectangular chips of
sizes `1 x 2` and `2 x 1`. Each chip fully occupies exactly two neighboring
squares of the field. Those chips do not overlap, but they may touch each other.
Overall, the field has exactly 28 chips, equal to the number of dominoes in the
set.

Now Gennady wants to replace each chip with a domino so that a magic figure
appears as a result. Different chips should be replaced by different dominoes.

Determine the number of contests Gennady can win over at the given position of
the chips. You are also required to find one possible way of replacing chips with
dominoes to win the next Codeforces round.

## Input

The first line contains two positive integers `n` and `m`
(`1 <= n, m <= 30`).

Each of the following `n` lines contains `m` characters describing the position
of chips on the field. Dots stand for empty cells. Latin letters from `a` to `z`
and `A`, `B` mark the chip positions. The two cells covered by the same chip are
marked by the same letter, and different chips are marked by different letters.

There are exactly 28 chips on the field. It is guaranteed that the field
description is correct.

It is also guaranteed that at least one solution exists.

## Output

Print on the first line the number of ways to replace chips with dominoes to get
a magic figure. This is the total number of contests that can be won using this
arrangement of chips.

Then print `n` lines containing `m` characters each: a field consisting of dots
and digits from `0` to `6`, representing any possible solution. All dominoes
should be different.

## Example

### Input

```text
8 8
.aabbcc.
.defghi.
kdefghij
klmnopqj
.lmnopq.
.rstuvw.
xrstuvwy
xzzAABBy
```

### Output

```text
10080
.001122.
.001122.
33440055
33440055
.225566.
.225566.
66113344
66113344
```

### editorial

First, let's understand that the entire field can be divided into 2x2 squares uniquely, and this can be done using a greedy algorithm. Now we have 14 squares. Next, we can make an optional transformation, which has made my life easier: let's construct a graph on 14 vertices. An edge between vertices will exist if and only if the corresponding squares share a domino. In principle, this is even a kind of time optimization, but that's not the point.
Now the problem is this: we have a graph on 14 vertices. We need to color them with colors from 0 to 6 such that all edges are different (two edges are the same if the endpoints have the same colors). It is claimed that the solution is an exhaustive search. We try the color of the first, second, third, ..., 14th vertex, and at each step we check that we haven't created an edge with the same coloring as an existing one (naturally, we do this with an array of bools). After that, all that remains is to print the answer and some coloring.
However, this solution may not be time-efficient. There's one more powerful optimization left. Note that there's essentially no difference between the colors. This means the colors in the coloring can be rearranged arbitrarily. Now let's learn how to count the number of answers accurate to color permutations, and finally multiply it by 7!=5040. This is simple: just introduce the condition that the color of the next vertex either occurred earlier or comes immediately after the maximum used color. For example, after colors 0 1 2 1 0 2, only colors 0..3 can come.
To be confident that this solution will work, we can roughly estimate the number of answers. Each vertex is colored in one of seven colors, each color occurs exactly twice, and we're not taking into account various permutations. The result is that, even with pruning at the last step of the recursion, it clearly fits within the TL.