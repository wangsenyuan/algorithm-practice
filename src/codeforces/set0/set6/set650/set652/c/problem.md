# Problem C - Foe Pairs

## Problem Statement

You are given a permutation `p` of length `n`. Also you are given `m` foe pairs `(ai, bi)` where `1 ≤ ai, bi ≤ n` and `ai ≠ bi`.

Your task is to count the number of different intervals `(x, y)` where `1 ≤ x ≤ y ≤ n` that do not contain any foe pairs. So you shouldn't count intervals `(x, y)` that contain at least one foe pair in it (the positions and order of the values from the foe pair are not important).

## Example

Consider `p = [1, 3, 2, 4]` and foe pairs are `{(3, 2), (4, 2)}`:
- The interval `(1, 3)` is **incorrect** because it contains a foe pair `(3, 2)`
- The interval `(1, 4)` is **incorrect** because it contains two foe pairs `(3, 2)` and `(4, 2)`
- The interval `(1, 2)` is **correct** because it doesn't contain any foe pair

## Input

- The first line contains two integers `n` and `m` (`1 ≤ n, m ≤ 3·10^5`) — the length of the permutation `p` and the number of foe pairs
- The second line contains `n` distinct integers `pi` (`1 ≤ pi ≤ n`) — the elements of the permutation `p`
- Each of the next `m` lines contains two integers `(ai, bi)` (`1 ≤ ai, bi ≤ n`, `ai ≠ bi`) — the i-th foe pair. Note a foe pair can appear multiple times in the given list

## Output

Print the only integer `c` — the number of different intervals `(x, y)` that does not contain any foe pairs.

**Note:** The answer can be too large, so you should use 64-bit integer type to store it. In C++ you can use the `long long` integer type and in Java you can use `long` integer type.

## Examples

### Example 1
**Input:**
```
4 2
1 3 2 4
3 2
2 4
```

**Output:**
```
5
```

### Example 2
**Input:**
```
9 5
9 7 2 3 1 4 6 5 8
1 6
4 5
2 7
7 2
2 7
```

**Output:**
```
20
```

## Note

In the first example the intervals from the answer are `(1, 1)`, `(1, 2)`, `(2, 2)`, `(3, 3)` and `(4, 4)`.


### ideas
1. 这个题目有点费解
2. pair (a, b) 表示的是两个数字，还是两个位置？
3. 应该是数字，如果是位置，就没p啥事了
4. 所以对于区间(l...r) 如果在这个区间中，存在数字a, 还存在数字b，那么就是存在一个pair了
5. 把pair当作一条从a到b的边，那么如果一个区间[l...r] 中间，不存在一条从a到b的边
6. 那么这个区间就是好的
7. 考虑对于r来说，如果区间l...r是好的，那么l+1...r也是好的
8. 所以可以对r来说，找到最小的l，满足l...r中间，没有pair
9. 对于[l...r]来说，假设其中有一个点b，如果存在两个a1 < a2 (< b) 那么只需要保留a2就可以了
10. 所以，对于l。。。r来说，就是这个区间内，如果最大的a >= l了，那么就要增加l
11. got 