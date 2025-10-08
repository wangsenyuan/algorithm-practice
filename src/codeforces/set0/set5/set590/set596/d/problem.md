# Problem D - Falling Trees

## Problem Description

There are `n` trees located at various positions on a line. Tree `i` is located at position `xi`. All the given positions of the trees are distinct.

The trees are equal, i.e. each tree has height `h`. Due to the wind, when a tree is cut down, it either falls left with probability `p`, or falls right with probability `1 - p`. If a tree hits another tree while falling, that tree will fall in the same direction as the tree that hit it. A tree can hit another tree only if the distance between them is strictly less than `h`.

## Example

For example, imagine there are 4 trees located at positions 1, 3, 5 and 8, while `h = 3` and the tree at position 1 falls right. It hits the tree at position 3 and it starts to fall too. In its turn it hits the tree at position 5 and it also starts to fall. The distance between 8 and 5 is exactly 3, so the tree at position 8 will not fall.

As long as there are still trees standing, Wilbur will select either the leftmost standing tree with probability 0.5 or the rightmost standing tree with probability 0.5. Selected tree is then cut down. If there is only one tree remaining, Wilbur always selects it. As the ground is covered with grass, Wilbur wants to know the expected total length of the ground covered with fallen trees after he cuts them all down because he is concerned about his grass-eating cow friends. Please help Wilbur.

## Input

The first line of the input contains two integers, `n` (1 ≤ `n` ≤ 2000) and `h` (1 ≤ `h` ≤ 10^8) and a real number `p` (0 ≤ `p` ≤ 1), given with no more than six decimal places.

The second line of the input contains `n` integers, `x1, x2, ..., xn` (-10^8 ≤ `xi` ≤ 10^8) in no particular order.

## Output

Print a single real number — the expected total length of the ground covered by trees when they have all fallen down. Your answer will be considered correct if its absolute or relative error does not exceed 10^-6.

Namely: let's assume that your answer is `a`, and the answer of the jury is `b`. The checker program will consider your answer correct, if `|a - b| / max(1, |b|) ≤ 10^-6`.

## Examples

### Example 1

**Input:**
```
2 2 0.500000
1 2
```

**Output:**
```
3.250000000
```

### Example 2

**Input:**
```
4 3 0.4
4 3 1 2
```

**Output:**
```
6.631200000
```

## Explanation

Consider the first example, we have 2 trees with height 2.

There are 3 scenarios:

1. **Both trees fall left.** This can either happen with the right tree falling left first, which has probability 0.5 × 0.5 = 0.25 (also knocking down the left tree), or the left tree can fall left and then the right tree can fall left, which has probability 0.5 × 0.5 = 0.25. Total probability is 0.5.

2. **Both trees fall right.** This is analogous to (1), so the probability of this happening is 0.5.

3. **The left tree falls left and the right tree falls right.** This is the only remaining scenario so it must have probability 0.5.

Cases 1 and 2 lead to a total of 3 units of ground covered, while case 3 leads to a total of 4 units of ground covered. Thus, the expected value is 0.5 × 3 + 0.5 × 4 = 3.5.


## Ideas

1. `dp[l][r][0/1]` 表示对于区间l...r，且左端的树往右倒，或者右端的树往左倒时的期望值
2. `dp[l][r][0] = 0.5 * dp[l+1][r][0] * p[i]` 如果选择左端的树，且它倒向右边（但是如果它倒向左边怎么办呢？）
3. (Additional ideas can be added here)

## Solution
Let us solve this problem using dynamic programming.

First let us reindex the trees by sorting them by x-coordinate. Let f(i, j, b1, b2) where we would like to consider the problem of if we only have trees i... j standing where b1 = 1 indicates that tree i - 1 falls right and b1 = 0 if it falls left and b2 = 1 indicates that tree j + 1 falls right and b2 = 0 if it falls left.

We start with the case that Wilbur chooses the left tree and it falls right. The plan is to calculate the expected length in this scenario and multiply by the chance of this case occurring, which is . We can easily calculate what is the farthest right tree that falls as a result of this and call it wi.

Then if wi >  = j this means the entire segment falls, from which the length of the ground covered by trees in i... j can be calculated. However, be careful when b2 = 0, as there may be overlapping covered regions when the tree j falls right but the tree j + 1 falls left.

If only wi < j, then we just consider adding the length of ground covered by trees i... wi falling right and add to the value of the subproblem f(wi + 1, j, 1, b2).

There is another interesting case where Wilbur chooses the left tree and it falls left. In this case we calculate the expected length and multiply by the chance of this occurring, which is . The expected length of ground covered by the trees here is just the length contributed by tree i falling left, which we must be careful calculating as there might be overlapping covered regions with the ith tree falling left and the i - 1th tree falling right. Then we also add the value of subproblem f(i + 1, j, 0, b2).

Doing this naively would take `O(n³)` time, but this can be lowered to `O(n²)` by precalculating what happens when tree `i` falls left or right.

We should also consider the cases that Wilbur chooses the right tree, but these cases are analogous by symmetry.

**Complexity:** `O(n²)`

