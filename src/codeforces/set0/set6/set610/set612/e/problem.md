# Problem E: Permutation Square Root

## Problem Description

A permutation of length n is an array containing each integer from 1 to n exactly once. For example, q = [4, 5, 1, 2, 3] is a permutation.

For the permutation q, the **square of permutation** is the permutation p such that p[i] = q[q[i]] for each i = 1...n. For example, the square of q = [4, 5, 1, 2, 3] is p = q² = [2, 3, 4, 5, 1].

This problem is about the **inverse operation**: given the permutation p, your task is to find such permutation q that q² = p. If there are several such q, find any of them.

## Input

The first line contains integer n (1 ≤ n ≤ 10⁶) — the number of elements in permutation p.

The second line contains n distinct integers p₁, p₂, ..., pₙ (1 ≤ pᵢ ≤ n) — the elements of permutation p.

## Output

- If there is no permutation q such that q² = p, print the number "-1".
- If the answer exists, print it. The only line should contain n different integers qᵢ (1 ≤ qᵢ ≤ n) — the elements of the permutation q. If there are several solutions, print any of them.

## Examples

### Example 1
**Input:**
```
4
2 1 4 3
```
**Output:**
```
3 4 2 1
```

### Example 2
**Input:**
```
4
2 1 3 4
```
**Output:**
```
-1
```

### Example 3
**Input:**
```
5
2 3 4 5 1
```
**Output:**
```
4 5 1 2 3
```


### ideas
1. 又是个让脑子抽的题目～
2. 假设 q = [... a.......]
3. q[i] = a,
4. q[q[i]] = q[a] = 在q中原来位置a处的数
5. q = [4, 5, 1, 2, 3] , p = [q[4], q[5], q[1], q[2], q[3]] = [2, 3, 4, 5, 1]
6. 1 -> 4 -> 2 -> 5 -> 3 -> 1
7. q = [3 4 2 1], 组成一个环, (1 -> 3 -> 2 -> 4 -> 1)
8. p = [q[3], q[4], q[2], q[1]] = [2, 1, 4, 3]
9. 这里感觉是把q中每个元素的下一个放置到对应p的位置
10. [2, 3, 1] => (3, 1, 2) 
11. a -> b -> c -> d -> e -> a
12. [c -> e -> b -> e -> a] （长度为5的没法产生2/3）
13. [a -> b -> c -> d] (b -> d -> b, c ->  a -> c) (ok)
14. [a -> b -> c -> d -> e -> f -> a]
15. 两个相等的，可以合并成一个更大的圈
16. 一个奇数的圈，还是自己
17. 偶数的圈，必须合并，奇数的圈保留