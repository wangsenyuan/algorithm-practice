# Problem D - Water Tanks Game

## Problem Statement

Mashmokh is playing a new game. In the beginning he has k liters of water and p coins. Additionally he has a rooted tree (an undirected connected acyclic graph) that consists of m vertices. Each vertex of the tree contains a water tank that is empty in the beginning.

The game begins with the fact that Mashmokh chooses some (no more than k) of these tanks (except the root) and pours into each of them exactly 1 liter of water. Then the following process is performed until there is no water remained in tanks.

## Game Process

The process consists of several steps:

1. At the beginning of each step Mashmokh opens doors of all tanks.
2. Then Mashmokh closes doors of some tanks (he is not allowed to close door of tank in the root) for the duration of this move. Let's denote the number of liters in some tank with closed door as w, Mashmokh pays w coins for the closing of that tank during this move.
3. Let's denote by x₁, x₂, ..., xₘ as the list of vertices of the tree sorted (nondecreasing) by their depth. The vertices from this list should be considered one by one in the order:
   - Firstly vertex x₁ (which is the root itself) is emptied.
   - Then for each vertex xᵢ (i > 1), if its door is closed then skip the vertex else move all the water from the tank of vertex xᵢ to the tank of its father (even if the tank of the father is closed).

## Scoring

Suppose l moves were made until the tree became empty. Let's denote the amount of water inside the tank of the root after the i-th move by wᵢ then Mashmokh will win max(w₁, w₂, ..., wₗ) dollars. Mashmokh wanted to know what is the maximum amount of dollars he can win by playing the above game. He asked you to find this value for him.

## Input

The first line of the input contains three space-separated integers m, k, p (2 ≤ m ≤ 10⁵; 0 ≤ k, p ≤ 10⁹).

Each of the following m - 1 lines contains two space-separated integers aᵢ, bᵢ (1 ≤ aᵢ, bᵢ ≤ m; aᵢ ≠ bᵢ) — the edges of the tree.

Consider that the vertices of the tree are numbered from 1 to m. The root of the tree has number 1.

## Output

Output a single integer, the number Mashmokh asked you to find.

## Examples

### Example 1

**Input:**
```
10 2 1
1 2
1 3
3 4
3 5
2 6
6 8
6 7
9 8
8 10
```

**Output:**
```
2
```

### Example 2

**Input:**
```
5 1000 1000
1 2
1 3
3 4
3 5
```

**Output:**
```
4
```


### ideas
1. 这个描述有好难理解～
2. 题目里面的door是个啥？ 所以这个图是翻过来的？root就是最低的那个点？
3. 所以就是在使用不超过p的token的前提下，尽可能的在root收集到足够的水
4. 而且w1应该是最大的。如果wi达到最大，可以在降低投入的情况下，w[1...i-1]都变成0
5. 且ans <= min(n-1, k)
6. 具体的花费是要阻拦的量。
7. 一个直观的感觉是，这些tank，尽量的在同一层上，这样的好处是，它们基本上可以同时到达root
8. 而且应该是底层（靠近root的地方）去拦截才有意义，而且要连续的一段
9. 而且放入的量，应该需要同时到达root，也就是这几层最终都要流到同一个层才可以
10. 除了最后一层，前面的所有层在放行前，都必须拦在最底层
11. 所以，应该是一段连续的， a1, a2, a3，ak
12. 且 $a_1 + sum(a_1 + a_2) + ... + sum(a_1 + a_2 + ... a_{k-1}) + a_k <= p$
13. 且这里每个$a_i$都有个上限，就是这层的节点数量
14. 所以确实是个双指针
15. 得倒过来处理，因为最后一层最好是填满的，这样花费比较少
16. $a_k + a_{k-1} * 2 + a_{k-2} * 3 + ... a2 * {k-1} + a1 * k$
17. 这里 k....2都是上限， a1是计算出来的 a1 >= 0, a1 <= c1
18. 所以在移动的时候， 减去suf[2...k], 然后加上suf[1..k-1]
19. 好麻烦～
20. a * 4 + b * 3 + c * 2 + d
21. 